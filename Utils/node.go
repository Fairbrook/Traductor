package Utils

type Segment struct {
	Lexema string `json:"Lexema"`
	Type   int    `json:"Type"`
	Index  int    `json:"Index"`
}

type Node struct {
	Segment Segment `json:"Segment"`
	State   int     `json:"State"`
}
