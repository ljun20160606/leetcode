package main

type problem struct {
	QuestionID                string        `json:"questionId"`
	QuestionFrontendID        string        `json:"questionFrontendId"`
	QuestionTitle             string        `json:"questionTitle"`
	QuestionTitleSlug         string        `json:"questionTitleSlug"`
	Content                   string        `json:"content"`
	Difficulty                string        `json:"difficulty"`
	Stats                     string        `json:"stats"`
	Contributors              string        `json:"contributors"`
	SimilarQuestions          string        `json:"similarQuestions"`
	DiscussURL                string        `json:"discussUrl"`
	MysqlSchemas              []interface{} `json:"mysqlSchemas"`
	RandomQuestionURL         string        `json:"randomQuestionUrl"`
	SessionID                 string        `json:"sessionId"`
	CategoryTitle             string        `json:"categoryTitle"`
	SubmitURL                 string        `json:"submitUrl"`
	InterpretURL              string        `json:"interpretUrl"`
	CodeDefinition            string        `json:"codeDefinitions"`
	SampleTestCase            string        `json:"sampleTestCase"`
	EnableTestMode            bool          `json:"enableTestMode"`
	MetaData                  string        `json:"metaData"`
	EnableRunCode             bool          `json:"enableRunCode"`
	EnableSubmit              bool          `json:"enableSubmit"`
	JudgerAvailable           bool          `json:"judgerAvailable"`
	InfoVerified              bool          `json:"infoVerified"`
	EnvInfo                   string        `json:"envInfo"`
	URLManager                string        `json:"urlManager"`
	Article                   string        `json:"article"`
	QuestionDetailURL         string        `json:"questionDetailUrl"`
	DiscussCategoryID         string        `json:"discussCategoryId"`
	DiscussSolutionCategoryID string        `json:"discussSolutionCategoryId"`
	LibraryURL                interface{}   `json:"libraryUrl"`
	CompanyTags               interface{}   `json:"companyTags"`
	TopicTags                 []struct {
		Name           string      `json:"name"`
		Slug           string      `json:"slug"`
		TranslatedName interface{} `json:"translatedName"`
	} `json:"topicTags"`
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

type codeDefinition struct {
	Text        string `json:"text"`
	Value       string `json:"value"`
	DefaultCode string `json:"defaultCode"`
}
