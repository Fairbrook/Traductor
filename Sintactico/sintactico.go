package Sintactico

import (
	"errors"
	"fmt"
	"strconv"

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

func reduce(rule Rule, next int, step *Step, pila *Utils.Stack) Utils.Tree {
	node := Utils.Node{
		Segment: Utils.Segment{
			Lexema: rule.terminal,
			Type:   rule.lexema,
		},
	}
	tree := Utils.Tree{}

	step.Output = "r" + strconv.FormatInt(int64(next*-1), 10)
	for i := 0; i < rule.popNumber; i++ {
		tree.AddChild(pila.Pop())
	}

	top := pila.Top.Root.State
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

func ProcessString(str string) (tree Utils.Tree, res []Step, err error) {
	var pila Utils.Stack
	var segment Lexico.Segment
	analizadorLexico := Lexico.Lexico{Input: str + "$"}
	node := initialNode()
	pila.Push(Utils.Tree{Root: node})
	for !pila.IsEmpty() {
		step := Step{Stack: pila.ToStr(), Input: analizadorLexico.GetLast()}
		segment, err = analizadorLexico.NextSegment()
		if err != nil {
			return
		}
		top := pila.Top.Root.State
		if next, ok := States[top][int(segment.State)]; ok {
			if next == 0 {
				err = errors.New(fmt.Sprintf("Cadena inesperada %s", segment.Lexema))
				return
			}

			if next >= 0 {
				strState := strconv.FormatInt(int64(next), 10)
				step.Output = "d" + strState
				node = segmentToNode(segment, next)
				pila.Push(Utils.Tree{Root: node})
				res = append(res, step)
				continue
			}

			analizadorLexico.GoBack()
			rule := Rules[next]

			if rule.lexema == accept {
				step.Output = "Aceptación"
				res = append(res, step)
				return
			}

			tree = reduce(rule, next, &step, &pila)
			pila.Push(tree)
			res = append(res, step)
			continue
		}
		err = errors.New(fmt.Sprintf("Cadena inesperada %s", segment.Lexema))
		return
	}
	return
}
