package main

import (
	"bytes"
	"fmt"
	"github.com/kataras/iris/core/errors"
	"github.com/ljun20160606/simplehttp"
	"gopkg.in/go-playground/pool.v3"
	"io"
	"net/url"
)

const (
	descriptionTemplate = "https://leetcode.com/problems/%v/description/"
	getQuestionDetail   = "getQuestionDetail"
)

// 不推荐这么写
var xCsrftoken string

func genAllFiles(stats []*stat) {
	p := pool.NewLimited(10)
	defer p.Close()
	batch := p.Batch()
	for e := range stats {
		batch.Queue(func(wu pool.WorkUnit) (r interface{}, err error) {
			stat := stats[e]
			descriptionUrl := fmt.Sprintf(descriptionTemplate, stat.QuestionTitleSlug)
			logger.Println("Fetch", stat.QuestionID, stat.QuestionTitle, descriptionUrl)
			err = simplehttp.NewRequest(defaultHttp2Client).Get().SetUrl(descriptionUrl).Send().Error()
			if err != nil {
				logger.Fatal(err)
				return
			}

			// 获取csrftoken，查询graphQL需要在header中添加x-csrftoken
			httpClient := defaultHttp2Client.(*simplehttp.HttpClient)
			parse, _ := url.Parse("https://leetcode.com")
			jar := httpClient.Jar.Cookies(parse)
			for _, v := range jar {
				if v.Name == "csrftoken" {
					xCsrftoken = v.Value
					break
				}
			}
			if xCsrftoken == "" {
				err = errors.New("x-csrftoken not in cookies")
				logger.Fatal(err)
				return
			}

			// 查询question
			qlResp, err := queryGraphQL(getQuestionDetail, stat.QuestionTitleSlug)
			if err != nil {
				logger.Fatal(err)
				return
			}
			q := newQuestionGenerator(&qlResp.Data.Question, "algorithms")
			// 生成文件夹
			q.genDir()
			// 生成题解
			q.genSolution()
			// 生成题目
			q.genQuestion()
			return
		})
	}
	batch.QueueComplete()
	batch.WaitAll()
}

// 查询题目
func queryGraphQL(operationName, titleSlug string) (*graphQLResp, error) {
	request := simplehttp.NewRequest(defaultHttp2Client)

	resp := request.Post().SetUrl("https://leetcode.com/graphql").SetBody(graphQLReader(operationName, titleSlug)).
		Head("referer", "https://leetcode.com/problems/two-sum/description/").
		Head("x-csrftoken", xCsrftoken).
		Head("content-type", "application/json").
		Send()
	qlResponse := new(graphQLResp)
	err := resp.JSON(qlResponse)
	if err != nil {
		return nil, err
	}
	return qlResponse, nil
}

// 构造graphQL的Body
func graphQLReader(operationName, titleSlug string) io.Reader {
	buffer := bytes.Buffer{}
	buffer.WriteString(`{
	"operationName": "` + operationName + `",
	"variables": {
		"titleSlug": "` + titleSlug + `"
	},
	"query": "query getQuestionDetail($titleSlug: String!) {\n  isCurrentUserAuthenticated\n  question(titleSlug: $titleSlug) {\n    questionId\n    questionFrontendId\n    questionTitle\n    questionTitleSlug\n    content\n    difficulty\n    stats\n    contributors\n    similarQuestions\n    discussUrl\n    mysqlSchemas\n    randomQuestionUrl\n    sessionId\n    categoryTitle\n    submitUrl\n    interpretUrl\n    codeDefinition\n    sampleTestCase\n    enableTestMode\n    metaData\n    enableRunCode\n    enableSubmit\n    judgerAvailable\n    infoVerified\n    envInfo\n    urlManager\n    article\n    questionDetailUrl\n    discussCategoryId\n    discussSolutionCategoryId\n    libraryUrl\n    companyTags {\n      name\n      slug\n      translatedName\n    }\n    topicTags {\n      name\n      slug\n      translatedName\n    }\n  }\n  interviewed {\n    interviewedUrl\n    companies {\n      id\n      name\n    }\n    timeOptions {\n      id\n      name\n    }\n    stageOptions {\n      id\n      name\n    }\n  }\n  subscribeUrl\n  isPremium\n  loginUrl\n}\n"
}`)
	return &buffer
}
