package main

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/ljun20160606/gox/fs"
	"io"
	"path/filepath"
	"strings"
	"text/template"
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
			QuestionTitle: q.Title,
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
	return fs.WriteFileNotExist(filepath.Join(q.DirPath, q.QuestionFrontendID+". "+q.Title+".go"), func(writer io.Writer) error {
		tpl, err := template.ParseFiles(problemTpl)
		if err != nil {
			return err
		}
		languages := []Language{Go, Java}
		snippet := getCodeSnippet(q.CodeSnippets, languages)
		if snippet == nil {
			var hint strings.Builder
			for i := range languages {
				hint.WriteString(languages[i].Lang)
				if i != len(languages)-1 {
					hint.WriteString(" or ")
				}
			}
			return errors.New("Not support " + hint.String() + ".")
		}
		err = tpl.Execute(writer, snippet)
		if err != nil {
			return err
		}
		return nil
	})
}

func getCodeSnippet(cs []codeSnippet, expect []Language) *codeSnippet {
	for i := range expect {
		for j := range cs {
			if cs[j].Lang == expect[i].Lang {
				return &cs[j]
			}
		}
	}
	return nil
}
