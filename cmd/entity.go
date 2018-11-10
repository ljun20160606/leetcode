package main

type graphQLResp struct {
	Data struct {
		Question problem `json:"question"`
	} `json:"data"`
}

type solutionTemplate struct {
	QuestionTitle string
	Content       string
}
