package Utils

type Segment struct {
	Lexema string
	Type   int
	Index  int
}

type Node struct {
	Segment Segment
	State   int
}
