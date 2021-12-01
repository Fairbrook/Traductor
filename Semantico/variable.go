package Semantico

import (
	"github.com/Fairbrook/analizador/Utils"
)

func defVar(tree *Utils.Tree, table *Table, pastType string) ([]error, []Symbol) {
	if tree.Children == nil {
		return make([]error, 0), make([]Symbol, 0)
	}
	variable := Variable{}
	iterator := tree.Children
	variable.Type = pastType
	if pastType == "" {
		variable.Type = iterator.Root.(Utils.Node).Segment.Lexema
	}

	iterator = iterator.Next
	variable.Segment = iterator.Root.(Utils.Node).Segment
	variable.Identifier = iterator.Root.(Utils.Node).Segment.Lexema

	localErrors, simbols := defVar(iterator.Next, table, variable.Type)
	// iterator = iterator.Next
	if table.Includes(variable.Identifier, true) {
		localErrors = append(localErrors, &Utils.SegmentError{
			Segment: variable.Segment,
			Message: Utils.DeclaredMsg,
		})
	}
	table.Set(&variable, nil)
	simbols = append(simbols, &variable)
	return localErrors, simbols
}
