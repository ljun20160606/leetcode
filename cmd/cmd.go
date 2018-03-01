package main

import (
	"errors"
	"github.com/urfave/cli"
	"gopkg.in/AlecAivazis/survey.v1"
	"strconv"
	"strings"
)

const (
	flagPath     = "path, p"
	flagForce    = "force, f"
	flagInteract = "interact, i"
)

var cmdLeetcode = cli.Command{
	Name:  "update",
	Usage: "抓取所有leetcode题目",
	Action: func(c *cli.Context) error {
		var path string
		if c.NArg() > 0 {
			path = c.Args()[0]
		} else {
			path = defaultLeetcodeJson
		}
		logger.Println("拉取所有leetcode题目至", path)
		genProblemsFile(path)
		return nil
	},
}

var cmdQuestion = cli.Command{
	Name:  "pull",
	Usage: "查询本地leetcode.json，然后抓取制定题目",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  flagForce,
			Usage: "先update再从表中查询",
		},
		cli.StringFlag{
			Name:  flagPath,
			Usage: "本地leetcode.json路径",
			Value: defaultLeetcodeJson,
		},
		cli.BoolFlag{
			Name:  flagInteract,
			Usage: "开启交互模式输入questionID",
		},
	},
	Action: func(c *cli.Context) error {
		var ids []int
		var idsStrs []string
		// 交互模式
		if c.Bool(flag(flagInteract)) {
			var idsStr = ""
			prompt := &survey.Input{Message: "请输入QuestionID"}
			survey.AskOne(prompt, &idsStr, nil)
			idsStrs = strings.Split(strings.Trim(idsStr, " "), " ")
			if len(idsStrs) == 0 || len(idsStrs[0]) == 0 {
				return errors.New("请输入 e.g. 1 2 3 4 5")
			}
		} else {
			// 直接输入
			if c.NArg() < 1 {
				return errors.New("pull [id...]")
			}
			idsStrs = c.Args()
		}
		// 初始化problems
		mustLoadProblemsFile(c.String(flag(flagPath)), c.Bool(flag(flagForce)))
		for e := range idsStrs {
			id, err := strconv.Atoi(idsStrs[e])
			if err != nil {
				return err
			}
			ids = append(ids, id)
		}
		stats := lqs.HasQuestionID(ids...)
		// 开始生成
		genAllFiles(stats)
		return nil
	},
}

func flag(f string) string {
	return strings.Split(f, ",")[0]
}
