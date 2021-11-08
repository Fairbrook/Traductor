package Sintactico

import (
	"errors"
	"fmt"

	Lexico "github.com/Fairbrook/analizador/Lexico"
	Utils "github.com/Fairbrook/analizador/Utils"
)

var accept int = 100

type Rule struct {
	terminal  string
	popNumber int
	lexema    int
}

type Step struct {
	Stack  string `json:"stack"`
	Input  string `json:"input"`
	Output string `json:"output"`
}

func segmentToNode(segment Lexico.Segment, state int) Utils.Node {
	return Utils.Node{
		Segment: Utils.Segment{
			Lexema: segment.Lexema,
			Index:  segment.Index,
			Type:   int(segment.State),
		},
		State: state,
	}
}

func reduce(rule Rule, next int, pila *Utils.Stack) Utils.Tree {
	node := Utils.Node{
		Segment: Utils.Segment{
			Lexema: rule.terminal,
			Type:   rule.lexema,
		},
	}
	tree := Utils.Tree{}
	for i := 0; i < rule.popNumber; i++ {
		tree.AddChild(pila.Pop().(Utils.Tree))
	}

	top := pila.Top.(Utils.Tree).Root.(Utils.Node).State
	action := States[top][rule.lexema]
	node.State = action
	tree.Root = node
	return tree
}

func initialNode() Utils.Node {
	return Utils.Node{
		Segment: Utils.Segment{
			Lexema: "$",
		},
		State: 0,
	}
}

func GenerateSyntacticTree(str string) (tree Utils.Tree, err error) {
	var pila Utils.Stack
	var segment Lexico.Segment
	currentLine := 0
	analizadorLexico := Lexico.Lexico{Input: str + "$"}
	node := initialNode()
	pila.Push(Utils.Tree{Root: node})
	for !pila.IsEmpty() {
		segment, err = analizadorLexico.NextSegment(currentLine)
		currentLine = segment.Line
		if err != nil {
			return
		}
		top := pila.Top.(Utils.Tree).Root.(Utils.Node).State
		if next, ok := States[top][int(segment.State)]; ok {
			if next == 0 {
				err = errors.New(fmt.Sprintf("Cadena inesperada %s", segment.Lexema))
				return
			}

			if next >= 0 {
				node = segmentToNode(segment, next)
				pila.Push(Utils.Tree{Root: node})
				continue
			}

			analizadorLexico.GoBack()
			rule := Rules[next]

			if rule.lexema == accept {
				return
			}

			tree = reduce(rule, next, &pila)
			pila.Push(tree)
			continue
		}
		err = errors.New(fmt.Sprintf("Cadena inesperada %s", segment.Lexema))
		return
	}
	return
}
