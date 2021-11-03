package testgame

type Test struct {
	ID        string     `json:"id"`
	Header    header     `json:"header"`
	Questions []question `json:"questions"`
	Social    social     `json:"social"`
}
