package Utils

// import "strconv"

type Segment struct {
	Lexema string `json:"Lexema"`
	Type   int    `json:"Type"`
	Index  int    `json:"Index"`
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
