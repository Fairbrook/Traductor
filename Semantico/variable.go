package Semantico

import (
	// "errors"
	"fmt"

	"github.com/Fairbrook/analizador/Utils"
)

func defVar(tree *Utils.Tree, table *Table, pastType string) []error {
	if tree.Children == nil {
		return make([]error, 0)
	}
	localErrors := []error{}
	variable := Variable{}
	iterator := tree.Children
	variable.Type = pastType
	if pastType == "" {
		variable.Type = iterator.Root.(Utils.Node).Segment.Lexema
	}

	iterator = iterator.Next
	variable.Identifier = iterator.Root.(Utils.Node).Segment.Lexema

	localErrors = append(localErrors, defVar(iterator.Next, table, variable.Type)...)
	// iterator = iterator.Next
	if table.Includes(variable.Identifier, true) {
		localErrors = append(localErrors, fmt.Errorf("el símbolo %s ya se encuentra declarado", variable.Identifier))
	}
	table.Set(&variable, nil)
	return localErrors
}
