package Utils

type Tree struct {
	Root     Node   `json:"Root"`
	Children []Tree `json:"Children"`
	index    int
}

func (t *Tree) AddChild(tree Tree) {
	if t.Children == nil {
		t.Children = make([]Tree, 10)
	}
	if len(t.Children) <= t.index {
		temp := make([]Tree, t.index+10)
		copy(temp, t.Children)
		t.Children = temp
	}
	t.Children[t.index] = tree
	t.index++
}

func (t *Tree) RemoveChild(pos int) {
	if len(t.Children) <= pos || pos < 0 {
		return
	}
	t.Children[pos] = t.Children[len(t.Children)-1]
	t.index--
}

func (t *Tree) GetChildren() []Tree {
	return t.Children[0:t.index]
}
