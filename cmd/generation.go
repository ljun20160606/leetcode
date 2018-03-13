package main

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/kataras/iris/core/errors"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unsafe"
)

func newQuestionGenerator(question *question, dirPrefix string) *questionGenerator {
	return &questionGenerator{
		question: question,
		DirPath:  filepath.Join(dirPrefix, offsetNumber(question.QuestionID, '0', 3)),
	}
}

type questionGenerator struct {
	*question
	DirPath string
}

// 生成题目文件
func (q *questionGenerator) genDir() {
	fileInfo, err := os.Stat(q.DirPath)
	if err != nil {
		err = os.MkdirAll(q.DirPath, 0755)
		if err != nil {
			panic(err)
		}
		return
	}
	if fileInfo.IsDir() {
		return
	}
	err = os.MkdirAll(q.DirPath, 0755)
	if err != nil {
		panic(err)
	}
}

func genINeed(name string, writeFunc func(writer io.Writer) error) error {
	_, err := os.Stat(name)
	if err != nil {
		file, err := os.Create(name)
		if err != nil {
			return err
		}
		defer file.Close()
		return writeFunc(file)
	}
	return nil
}

// 生成readme
func (q *questionGenerator) genSolution() {
	fileP := filepath.Join(q.DirPath, "README.md")
	need := genINeed(fileP, func(writer io.Writer) error {
		tpl, err := template.ParseFiles("cmd/solution.tpl.txt")
		if err != nil {
			return err
		}
		document, err := goquery.NewDocumentFromReader(strings.NewReader(q.Content))
		if err != nil {
			return err
		}
		m := make(map[string]string)
		m["QuestionTitle"] = q.QuestionTitle
		m["Content"] = document.Text()
		err = tpl.Execute(writer, m)
		if err != nil {
			return err
		}
		return nil
	})
	if need != nil {
		logger.Fatal(need)
		panic(need)
	}
}

// 生成题目question
func (q *questionGenerator) genQuestion() {
	fileP := filepath.Join(q.DirPath, q.QuestionFrontendID+". "+q.QuestionTitle+".go")
	need := genINeed(fileP, func(writer io.Writer) error {
		tpl, err := template.ParseFiles("cmd/question.tpl.txt")
		if err != nil {
			return err
		}
		var cd codeDefinition
		bytes := *(*[]byte)(unsafe.Pointer(&q.CodeDefinition))
		err = json.Unmarshal(bytes, &cd)
		if err != nil {
			return err
		}
		for i := range cd {
			if cd[i].Value == "golang" {
				err = tpl.Execute(writer, cd[i])
				if err != nil {
					return err
				}
				return nil
			}
		}
		return errors.New("Not support golang.")
	})
	if need != nil {
		logger.Fatal(need)
		panic(need)
	}
}

// 填补 e.g. 1 -> 001
func offsetNumber(c string, offsetByte byte, amount int) string {
	if len(c) > amount {
		return c
	}
	missing := amount - len(c)
	builder := strings.Builder{}
	for i := 0; i < missing; i++ {
		builder.Write([]byte{offsetByte})
	}
	builder.WriteString(c)
	return builder.String()
}
