package Utils

type Stack struct {
	Top     string
	content []string
	index   int
}

func (q *Stack) Push(str string) {
	q.Top = str
	if q.content == nil {
		q.content = make([]string, 10)
		q.index = -1
	}
	q.index++
	if len(q.content) <= q.index {
		temp := make([]string, q.index+10)
		copy(temp, q.content)
		q.content = temp
	}
	q.content[q.index] = str
}

func (q *Stack) Pop() string {
	if q.content == nil || q.index < 0 {
		return ""
	}
	q.index--
	if len(q.content)-q.index > 10 {
		q.content = q.content[0 : q.index+10]
	}
	q.Top = q.content[q.index]
	return q.content[q.index+1]
}

func (q *Stack) ToStr() string {
	str := ""
	for i := 0; i <= q.index; i++ {
		str += q.content[i]
		if i < q.index {
			str += " "
		}
	}
	return str
}

func (q *Stack) IsEmpty() bool {
	return len(q.content) == 0 || q.index < 0
}
