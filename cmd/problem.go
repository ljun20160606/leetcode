package main

type problem struct {
	QuestionID            string        `json:"questionId"`
	QuestionFrontendID    string        `json:"questionFrontendId"`
	BoundTopicID          int           `json:"boundTopicId"`
	Title                 string        `json:"title"`
	TitleSlug             string        `json:"titleSlug"`
	Content               string        `json:"content"`
	TranslatedTitle       string        `json:"translatedTitle"`
	TranslatedContent     string        `json:"translatedContent"`
	IsPaidOnly            bool          `json:"isPaidOnly"`
	Difficulty            string        `json:"difficulty"`
	Likes                 int           `json:"likes"`
	Dislikes              int           `json:"dislikes"`
	IsLiked               interface{}   `json:"isLiked"`
	SimilarQuestions      string        `json:"similarQuestions"`
	Contributors          []interface{} `json:"contributors"`
	LangToValidPlayground string        `json:"langToValidPlayground"`
	TopicTags             []struct {
		Name           string `json:"name"`
		Slug           string `json:"slug"`
		TranslatedName string `json:"translatedName"`
		Typename       string `json:"__typename"`
	} `json:"topicTags"`
	CompanyTagStats interface{}   `json:"companyTagStats"`
	CodeSnippets    []codeSnippet `json:"codeSnippets"`
	Stats           string        `json:"stats"`
	Hints           []interface{} `json:"hints"`
	Solution        interface{}   `json:"solution"`
	Status          interface{}   `json:"status"`
	SampleTestCase  string        `json:"sampleTestCase"`
	MetaData        string        `json:"metaData"`
	JudgerAvailable bool          `json:"judgerAvailable"`
	JudgeType       string        `json:"judgeType"`
	MysqlSchemas    []interface{} `json:"mysqlSchemas"`
	EnableRunCode   bool          `json:"enableRunCode"`
	EnableTestMode  bool          `json:"enableTestMode"`
	EnvInfo         string        `json:"envInfo"`
	Typename        string        `json:"__typename"`
}

type stat struct {
	TotalAcs            int    `json:"total_acs"`
	QuestionTitle       string `json:"question__title"`
	IsNewQuestion       bool   `json:"is_new_question"`
	QuestionArticleSlug string `json:"question__article__slug"`
	TotalSubmitted      int    `json:"total_submitted"`
	FrontendQuestionID  int    `json:"frontend_question_id"`
	QuestionTitleSlug   string `json:"question__title_slug"`
	QuestionArticleLive bool   `json:"question__article__live"`
	QuestionHide        bool   `json:"question__hide"`
	QuestionID          int    `json:"question_id"`
}

type codeSnippet struct {
	Lang     string `json:"lang"`
	LangSlug string `json:"langSlug"`
	Code     string `json:"code"`
	Typename string `json:"__typename"`
}
