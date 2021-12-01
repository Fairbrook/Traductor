package Utils

// import "strconv"

type Segment struct {
	Lexema    string `json:"lexema"`
	State     uint8  `json:"state"`
	StateName string `json:"state_name"`
	Index     int    `json:"index"`
	Line      int    `json:"line"`
	Type      int    `json:"type"`
}

type Node struct {
	Segment Segment `json:"Segment"`
	State   int     `json:"State"`
}

// func (n *Node) ToStr() string {
// 	str := ""
// 	str += n.Segment.Lexema + " " + strconv.FormatInt(int64(n.State), 10)
// 	return str
// }
