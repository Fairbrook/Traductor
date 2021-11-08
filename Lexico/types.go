package Lexico

type Segment struct {
	Lexema    string `json:"lexema"`
	State     uint8  `json:"state"`
	StateName string `json:"state_name"`
	Index     int    `json:"index"`
	Line      int    `json:"line"`
}
