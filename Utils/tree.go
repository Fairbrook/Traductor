package Utils

type Tree struct {
	Root     interface{} `json:"Root"`
	Children *Tree       `json:"Children"`
	Current  *Tree
	Next     *Tree
	Prev     *Tree
	Length   int
}

func (t *Tree) AddChild(tree Tree) {
	t.Length++
	if t.Children == nil {
		t.Children = new(Tree)
		*t.Children = tree
		t.Current = t.Children
		return
	}
	prev := t.Current

	t.Current = new(Tree)
	*t.Current = tree

	t.Children = t.Current
	t.Children.Next = prev
	prev.Prev = t.Current
}
