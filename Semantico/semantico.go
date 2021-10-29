package Semantico

import (
	// "errors"
	// "fmt"

	"github.com/Fairbrook/analizador/Sintactico"
	"github.com/Fairbrook/analizador/Utils"
)

func def(tree *Utils.Tree, table *Table) []error {
	if tree.Children == nil {
		return make([]error, 0)
	}
	defType := tree.Root.(Utils.Node).Segment.Lexema
	switch defType {
	case "Definiciones":
		{
			localErrors := def(tree.Children, table)
			return append(localErrors, def(tree.Children.Next, table)...)
		}
	case "Definicion":
		return def(tree.Children, table)
	case "DefVar":
		return defVar(tree, table, "")
	case "DefFunc":
		return defFunc(tree, table)
	default:
		return make([]error, 0)

	}
}

func Analize(str string) (table Table, errs []error) {
	var err error
	var tree Utils.Tree
	table = Table{}
	if tree, err = Sintactico.GenerateSyntacticTree(str); err != nil {
		errs = []error{err}
		return
	}
	if tree.Children == nil {
		return
	}
	errs = def(tree.Children, &table)
	return
}
