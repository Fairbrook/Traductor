package Utils

import "strconv"

type Stack struct {
	Top     Tree
	content []Tree
	index   int
}

func (q *Stack) Push(tree Tree) {
	q.Top = tree
	if q.content == nil {
		q.content = make([]Tree, 10)
		q.index = -1
	}
	q.index++
	if len(q.content) <= q.index {
		temp := make([]Tree, q.index+10)
		copy(temp, q.content)
		q.content = temp
	}
	q.content[q.index] = tree
}

func (q *Stack) Pop() Tree {
	if q.content == nil || q.index < 0 {
		return Tree{}
	}
	q.index--
	if len(q.content)-q.index > 10 {
		q.content = q.content[0 : q.index+10]
	}
	if q.index >= 0 {
		q.Top = q.content[q.index]
	}
	return q.content[q.index+1]
}

func (q *Stack) ToStr() string {
	str := ""
	for i := 0; i <= q.index; i++ {
		str += q.content[i].Root.Segment.Lexema + " " + strconv.FormatInt(int64(q.content[i].Root.State), 10)
		if i < q.index {
			str += "\t"
		}
	}
	return str
}

func (q *Stack) IsEmpty() bool {
	return len(q.content) == 0 || q.index < 0
}
