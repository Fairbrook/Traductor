package Utils

type Tree struct {
	Root     Node
	children []Tree
	index    int
}

func (t *Tree) AddChild(tree Tree) {
	if t.children == nil {
		t.children = make([]Tree, 10)
	}
	if len(t.children) <= t.index {
		temp := make([]Tree, t.index+10)
		copy(temp, t.children)
		t.children = temp
	}
	t.children[t.index] = tree
}

func (t *Tree) RemoveChild(pos int) {
	if len(t.children) <= pos || pos < 0 {
		return
	}
	t.children[pos] = t.children[len(t.children)-1]
	t.index--
}

func (t *Tree) GetChildren() []Tree {
	return t.children[0:t.index]
}
