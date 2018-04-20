package main

import (
	"fmt"
	"github.com/kataras/iris/core/errors"
	"github.com/ljun20160606/di"
	"github.com/ljun20160606/simplehttp"
	"github.com/urfave/cli"
	"gopkg.in/go-playground/pool.v3"
	"io"
	"log"
	"net/url"
	"os"
	"sort"
	"strings"
	"sync"
	"text/template"
)

const (
	algorithms    = "algorithms"
	csrftoken     = "csrftoken"
	README        = "README.md"
	LeetcodeQuery = `query getQuestionDetail($titleSlug: String!) {
  isCurrentUserAuthenticated
  question(titleSlug: $titleSlug) {
    questionId
    questionFrontendId
    questionTitle
    translatedTitle
    questionTitleSlug
    content
    translatedContent
    difficulty
    stats
    contributors
    similarQuestions
    discussUrl
    mysqlSchemas
    randomQuestionUrl
    sessionId
    categoryTitle
    submitUrl
    interpretUrl
    codeDefinition
    sampleTestCase
    enableTestMode
    metaData
    enableRunCode
    enableSubmit
    judgerAvailable
    infoVerified
    envInfo
    urlManager
    article
    questionDetailUrl
    discussCategoryId
    discussSolutionCategoryId
    libraryUrl
    companyTags {
      name
      slug
      translatedName
    }
    topicTags {
      name
      slug
      translatedName
    }
  }
  interviewed {
    interviewedUrl
    companies {
      id
      name
      slug
    }
    timeOptions {
      id
      name
    }
    stageOptions {
      id
      name
    }
  }
  subscribeUrl
  isPremium
  loginUrl
}`
)

var (
	lc     = &leetcodeCli{}
	logger = log.New(os.Stderr, "[leetcode-cli] ", log.Lshortfile|log.LstdFlags)
)

func init() {
	di.Put(lc)
	di.Put(simplehttp.NewClient(&simplehttp.Builder{ProtoMajor: simplehttp.HTTP2}))
}

type leetcodeCli struct {
	xCsrftoken string
	problems   *problems
	lock       sync.Mutex
	items      []item
	Commands   []cli.Command
	HttpClient simplehttp.Client `di:"*"`
	Config     *config           `di:"*"`
}

func (l *leetcodeCli) Init() {
	l.problems = new(problems)
}

func (l *leetcodeCli) AddCommand(command cli.Command) {
	l.Commands = append(l.Commands, command)
}

func (l *leetcodeCli) AddItem(q *question) {
	var topics []string
	for e := range q.TopicTags {
		topics = append(topics, q.TopicTags[e].Name)
	}
	i := item{
		ID:         q.QuestionID,
		Title:      fmt.Sprintf("[%v](algorithms/%v/README.md)", q.QuestionTitle, q.FileName),
		Topics:     strings.Join(topics, ", "),
		Difficulty: q.Difficulty,
	}
	for e := range l.items {
		if l.items[e].ID == q.QuestionID {
			l.items[e] = i
			return
		}
	}
	l.lock.Lock()
	l.items = append(l.items, i)
	l.lock.Unlock()
}

func (l *leetcodeCli) SortItem() {
	sort.Sort(itemSlice(l.items))
}

// 根据questionID生成 {question}.go README.md
func (l *leetcodeCli) Gen(ids []int) {
	p := pool.NewLimited(10)
	defer p.Close()
	batch := p.Batch()
	stats := l.problems.hasQuestionID(ids...)
	for e := range stats {
		stat := stats[e]
		batch.Queue(func(wu pool.WorkUnit) (r interface{}, err error) {
			descriptionUrl := fmt.Sprintf(l.Config.DescriptionTemplate, stat.QuestionTitleSlug)
			logger.Println("Fetch", stat.QuestionID, stat.QuestionTitle, descriptionUrl)
			err = simplehttp.NewRequest(l.HttpClient).Get().SetUrl(descriptionUrl).Send().Error()
			if err != nil {
				logger.Fatal(err)
				return
			}

			// 获取csrftoken，查询graphQL需要在header中添加x-csrftoken
			httpClient := l.HttpClient.(*simplehttp.HttpClient)
			parse, _ := url.Parse(l.Config.LeetcodeCom)
			jar := httpClient.Jar.Cookies(parse)
			for _, v := range jar {
				if v.Name == csrftoken {
					l.xCsrftoken = v.Value
					break
				}
			}

			if l.xCsrftoken == "" {
				err = errors.New("x-csrftoken not in cookies")
				logger.Fatal(err)
				return
			}

			// 查询question
			qlResp, err := l.queryGraphQL(descriptionUrl, l.Config.OperationName, stat.QuestionTitleSlug)
			if err != nil {
				logger.Fatal(err)
				return
			}
			q := newQuestion(&qlResp.Data.Question, algorithms)

			// 生成文件夹
			err = fs.MkdirP(q.DirPath)
			if err != nil {
				return
			}

			// 生成题解
			err = q.WriteSolution(l.Config.SolutionTpl)
			if err != nil {
				return
			}

			// 生成题目 problem
			err = q.WriteProblem(l.Config.ProblemTpl)
			if err != nil {
				return
			}

			// readme添加item
			l.AddItem(q)
			return
		})
	}
	batch.QueueComplete()
	batch.WaitAll()
}

// 读取algorithms.json
func (l *leetcodeCli) ReadLock() {
	err := fs.ReadJSON(l.Config.AlgorithmsLock, &l.items)
	if err != nil {
		logger.Println(err)
		panic(err)
	}
}

// 写algorithms.json
func (l *leetcodeCli) WriteLock() {
	err := fs.WriteJSON(l.Config.AlgorithmsLock, &l.items)
	if err != nil {
		logger.Println(err)
		panic(err)
	}
}

// 查询题目
func (l *leetcodeCli) queryGraphQL(referer, operationName, titleSlug string) (*graphQLResp, error) {
	request := simplehttp.NewRequest(l.HttpClient)

	resp := request.Get().SetUrl(l.Config.Graphql).
		Head("referer", referer).
		Head("x-csrftoken", l.xCsrftoken).
		Head("content-type", "application/json, text/plain").
		Query("query", LeetcodeQuery).
		Query("operationName", operationName).
		Query("variables", `{"titleSlug":"`+titleSlug+`"}`).
		Send()
	qlResponse := new(graphQLResp)
	err := resp.JSON(qlResponse)
	if err != nil {
		return nil, err
	}
	return qlResponse, nil
}

// 读取leetcode.json
func (l *leetcodeCli) ReadCatalog(path string, reGen bool) {
	if !fs.Exists(path) || reGen {
		l.WriteCatalog(path)
	}
	err := fs.ReadJSON(path, l.problems)
	if err != nil {
		panic(err)
	}
}

// 生成leetcode.json
func (l *leetcodeCli) WriteCatalog(path string) {
	fs.WriteFileNotExist(path, func(writer io.Writer) error {
		b, err := simplehttp.NewRequest(l.HttpClient).Get().SetUrl(l.Config.ProblemsAllUrl).Send().Body()
		if err != nil {
			panic(err)
		}
		_, err = writer.Write(b)
		if err != nil {
			panic(err)
		}
		return nil
	})
}

func (l *leetcodeCli) WriteReadme() {
	fs.WriteFile(README, func(writer io.Writer) error {
		tpl, err := template.ParseFiles(l.Config.ReadmeTpl)
		if err != nil {
			return err
		}
		err = tpl.Execute(writer, l.items)
		if err != nil {
			return err
		}
		return nil
	})
}
