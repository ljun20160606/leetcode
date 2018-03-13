package main

type question struct {
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
	CodeDefinition            string        `json:"codeDefinition"`
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

type leetcodeQuestions struct {
	FrequencyMid    int    `json:"frequency_mid"`
	FrequencyHigh   int    `json:"frequency_high"`
	AcMedium        int    `json:"ac_medium"`
	AcEasy          int    `json:"ac_easy"`
	NumSolved       int    `json:"num_solved"`
	CategorySlug    string `json:"category_slug"`
	StatStatusPairs []struct {
		Status     interface{} `json:"status"`
		Stat       stat        `json:"stat"`
		IsFavor    bool        `json:"is_favor"`
		PaidOnly   bool        `json:"paid_only"`
		Difficulty struct {
			Level int `json:"level"`
		} `json:"difficulty"`
		Frequency int `json:"frequency"`
		Progress  int `json:"progress"`
	} `json:"stat_status_pairs"`
	AcHard   int    `json:"ac_hard"`
	UserName string `json:"user_name"`
	NumTotal int    `json:"num_total"`
}

type codeDefinition []struct {
	Text        string `json:"text"`
	Value       string `json:"value"`
	DefaultCode string `json:"defaultCode"`
}

func (l *leetcodeQuestions) HasQuestionID(ids ...int) []*stat {
	idMap := make(map[int]bool)
	for e := range ids {
		idMap[ids[e]] = false
	}
	counter := len(ids)
	var stats []*stat
	for i := range l.StatStatusPairs {
		if counter == 0 {
			return stats
		}
		v := l.StatStatusPairs[i]
		if vv, ok := idMap[v.Stat.QuestionID]; ok {
			if vv {
				continue
			}
			counter--
			idMap[v.Stat.QuestionID] = true
			stats = append(stats, &v.Stat)
		}
	}
	return stats
}

type graphQLResp struct {
	Data struct {
		IsCurrentUserAuthenticated bool     `json:"isCurrentUserAuthenticated"`
		Question                   question `json:"question"`
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
