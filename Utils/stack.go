package Utils

type StackNode struct {
	Data interface{}
	Next *StackNode
}

type Stack struct {
	Top    interface{}
	root   *StackNode
	Length int
}

func (q *Stack) Push(data interface{}) {
	q.Top = data
	q.Length++
	prev := q.root
	newRoot := new(StackNode)
	*newRoot = StackNode{
		Data: data,
		Next: nil,
	}
	if q.root != nil {
		newRoot.Next = prev
	}
	q.root = newRoot

}

func (q *Stack) Pop() interface{} {
	data := q.Top
	q.Length--
	if q.root == nil {
		q.Top = *new(interface{})
		return *new(interface{})
	}
	q.root = q.root.Next
	q.Top = q.root.Data
	return data
}

// func (q *Stack) ToStr() string {
// 	str := ""
// 	iterator := q.root
// 	for iterator != nil {
// 		str += iterator.Data.Root.ToStr()
// 		if iterator.Next != nil {
// 			str += "\t"
// 		}
// 		iterator = iterator.Next
// 	}
// 	return str
// }

func (q *Stack) IsEmpty() bool {
	return q.root == nil
}

func (q *Stack) GetListPointer() *StackNode {
	return q.root
}
