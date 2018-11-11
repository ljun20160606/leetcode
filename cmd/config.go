package main

import "github.com/ljun20160606/di"

func init() {
	di.Put(new(config))
}

type config struct {
	DefaultLeetcodeJson string `di:"#.defaultLeetcodeJson"`
	OperationName       string `di:"#.operationName"`
	LeetcodeCom         string `di:"#.leetcodeCom"`
	Graphql             string `di:"#.graphql"`
	ProblemsAllUrl      string `di:"#.problemsAllUrl"`
	DescriptionTemplate string `di:"#.descriptionTemplate"`
	SolutionTpl         string `di:"#.solutionTpl"`
	ProblemTpl          string `di:"#.problemTpl"`
	ReadmeTpl           string `di:"#.readmeTpl"`
	UnittestTpl         string `di:"#.unittestTpl"`
	AlgorithmsLock      string `di:"#.algorithmsLock"`
}
