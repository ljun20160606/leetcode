package main

import (
	"encoding/json"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/ljun20160606/gox/fs"
	"io"
	"path/filepath"
	"strings"
	"text/template"
	"unsafe"
)

type question struct {
	*problem
	FileName string
	DirPath  string
}

func newQuestion(problem *problem, dirPrefix string) *question {
	fileName := offsetNumber(problem.QuestionID, '0', 3)
	return &question{
		problem:  problem,
		FileName: fileName,
		DirPath:  filepath.Join(dirPrefix, fileName),
	}
}

func (q *question) WriteSolution(solutionTplPath string) error {
	return fs.WriteFileNotExist(filepath.Join(q.DirPath, README), func(writer io.Writer) error {
		tpl, err := template.ParseFiles(solutionTplPath)
		if err != nil {
			return err
		}
		document, err := goquery.NewDocumentFromReader(strings.NewReader(q.Content))
		if err != nil {
			return err
		}
		s := solutionTemplate{
			QuestionTitle: q.QuestionTitle,
			Content:       document.Text(),
		}
		err = tpl.Execute(writer, s)
		if err != nil {
			return err
		}
		return nil
	})
}

func (q *question) WriteProblem(problemTpl string) error {
	return fs.WriteFileNotExist(filepath.Join(q.DirPath, q.QuestionFrontendID+". "+q.QuestionTitle+".go"), func(writer io.Writer) error {
		tpl, err := template.ParseFiles(problemTpl)
		if err != nil {
			return err
		}
		var cs []codeDefinition
		b := *(*[]byte)(unsafe.Pointer(&q.CodeDefinition))
		err = json.Unmarshal(b, &cs)
		if err != nil {
			return err
		}
		expect := []string{"golang", "cpp"}
		definition := getCodeDefinition(cs, expect)
		if definition == nil {
			return errors.New("Not support " + strings.Join(expect, " or ") + ".")
		}
		err = tpl.Execute(writer, definition)
		if err != nil {
			return err
		}
		return nil
	})
}

func getCodeDefinition(cs []codeDefinition, expect []string) *codeDefinition {
	for i := range expect {
		for j := range cs {
			if cs[j].Value == expect[i] {
				return &cs[j]
			}
		}
	}
	return nil
}
