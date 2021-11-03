package testgame

type question struct {
	Text          string   `json:"text"`
	AnswerOptions []string `json:"answerOptions"`
	Answer        int      `json:"answer"`
}
