package main

type graphQLResp struct {
	Data struct {
		IsCurrentUserAuthenticated bool    `json:"isCurrentUserAuthenticated"`
		Question                   problem `json:"question"`
		Interviewed                struct {
			InterviewedURL string `json:"interviewedUrl"`
			Companies      []struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"companies"`
			TimeOptions []struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"timeOptions"`
			StageOptions []struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"stageOptions"`
		} `json:"interviewed"`
		SubscribeURL string `json:"subscribeUrl"`
		IsPremium    bool   `json:"isPremium"`
		LoginURL     string `json:"loginUrl"`
	} `json:"data"`
}

type solutionTemplate struct {
	QuestionTitle string
	Content       string
}
