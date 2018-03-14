package main

type problems struct {
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

// 获取存在的question
func (lq *problems) hasQuestionID(ids ...int) []*stat {
	idMap := make(map[int]bool)
	for e := range ids {
		idMap[ids[e]] = false
	}
	counter := len(ids)
	var stats []*stat
	for i := range lq.StatStatusPairs {
		if counter == 0 {
			return stats
		}
		v := lq.StatStatusPairs[i]
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
