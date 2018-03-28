package main

import (
	"encoding/json"
	"github.com/ljun20160606/simplehttp"
	"io/ioutil"
	"log"
	"os"
)

const (
	defaultLeetcodeJson = "leetcode.json"
)

var (
	logger             = log.New(os.Stderr, "[leetcode-cli] ", log.Lshortfile|log.LstdFlags)
	defaultHttp2Client = simplehttp.NewClient(&simplehttp.Builder{ProtoMajor: simplehttp.HTTP2})
	lqs                *leetcodeQuestions
)

// 加载leetcode字典
func mustLoadProblemsFile(path string, reGen bool) {
	if _, err := os.Stat(path); os.IsNotExist(err) || reGen {
		genProblemsFile(path)
	}
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var q leetcodeQuestions
	err = json.Unmarshal(bytes, &q)
	if err != nil {
		panic(err)
	}
	lqs = &q
}

// 生成leetcode题目字典
func genProblemsFile(path string) {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	s, err := simplehttp.NewRequest(defaultHttp2Client).Get().SetUrl(problemsAllUrl).Send().String()
	if err != nil {
		panic(err)
	}
	_, err = file.WriteString(s)
	if err != nil {
		panic(err)
	}
}
