package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

// 1. 生成leetcode.json
// 2. 根据题目生成题目所需文件
//   (1) 判断文件是否在字典中
//   (2) 抓取描述
//	 (3) 根据Problem建立文件
func main() {
	app := cli.NewApp()
	app.Name = "leetcode"
	app.Version = "0.0.1"
	app.Usage = "leetcode刷题辅助工具"
	app.Commands = append(app.Commands,
		cmdLeetcode,
		cmdQuestion)
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
