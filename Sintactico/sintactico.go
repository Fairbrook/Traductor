package Sintactico

import (
	"errors"
	"fmt"
	"strconv"

	Lexico "github.com/Fairbrook/analizador/Lexico"
	Utils "github.com/Fairbrook/analizador/Utils"
)

var E int = 103
var accept int = 104

var states = map[int]map[int]int{
	0: {0: 2, E: 1},
	1: {23: -1},
	2: {5: 3, 23: -3},
	3: {0: 2, E: 4},
	4: {23: -2},
}

var rules = map[int][3]int{
	-1: {accept, accept, ' '},
	-2: {E, 3, 'E'},
	-3: {E, 1, 'E'},
}

type Step struct {
	Stack  string `json:"stack"`
	Input  string `json:"input"`
	Output string `json:"output"`
}

func ProcessString(str string) (res []Step, err error) {
	var pila Utils.Stack
	var segment Lexico.Segment
	analizadorLexico := Lexico.Lexico{Input: str + "$"}
	node := Utils.Node{
		Segment: Utils.Segment{
			Lexema: "$",
		},
		State: 0,
	}
	pila.Push(Utils.Tree{Root: node})
	for !pila.IsEmpty() {
		step := Step{Stack: pila.ToStr(), Input: analizadorLexico.GetLast()}
		segment, err = analizadorLexico.NextSegment()
		if err != nil {
			return
		}
		top := pila.Top.Root.State
		if next, ok := states[top][int(segment.State)]; ok {
			if next >= 0 {
				strState := strconv.FormatInt(int64(next), 10)
				step.Output = "d" + strState
				node = Utils.Node{
					Segment: Utils.Segment{
						Lexema: segment.Lexema,
						Index:  segment.Index,
						Type:   int(segment.State),
					},
					State: next,
				}
				pila.Push(Utils.Tree{Root: node})
				res = append(res, step)
				continue
			}

			analizadorLexico.GoBack()
			rule := rules[next]

			if rule[0] == accept {
				step.Output = "Aceptación"
				res = append(res, step)
				return
			}

			node = Utils.Node{
				Segment: Utils.Segment{
					Lexema: string(byte(rule[2])),
					Type:   rule[1],
				},
			}
			tree := Utils.Tree{}

			step.Output = "r" + strconv.FormatInt(int64(next*-1), 10)
			for i := 0; i < rule[1]; i++ {
				tree.AddChild(pila.Pop())
			}

			top = pila.Top.Root.State
			action := states[top][rule[0]]
			node.State = action
			tree.Root = node
			pila.Push(tree)
			res = append(res, step)
			continue
		}
		err = errors.New(fmt.Sprintf("Cadena inesperada %s", segment.Lexema))
		return
	}
	return
}
