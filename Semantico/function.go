package Semantico

import (
	"fmt"

	"github.com/Fairbrook/analizador/Utils"
)

func localDef(tree *Utils.Tree, table *Table, function Function) []error {
	if tree.Children == nil {
		return make([]error, 0)
	}
	defType := tree.Root.(Utils.Node).Segment.Lexema
	switch defType {
	case "DefLocales":
		{
			localErrors := localDef(tree.Children, table, function)
			return append(localErrors, localDef(tree.Children.Next, table, function)...)
		}
	case "DefLocal":
		return localDef(tree.Children, table, function)
	case "DefVar":
		return defVar(tree, table, "")
	case "Sentencia":
		return procSentencia(tree, table, function)
	default:
		return make([]error, 0)

	}
}

func getParams(params *Utils.Tree, table *Table) (parameters map[string]Variable, errs []error) {
	if params.Children == nil {
		parameters = map[string]Variable{}
		return
	}
	iterator := params.Children
	lexema := iterator.Root.(Utils.Node).Segment.Lexema

	if lexema == "," {
		iterator = iterator.Next
	}

	variable := Variable{}
	variable.Type = iterator.Root.(Utils.Node).Segment.Lexema

	iterator = iterator.Next
	variable.Identifier = iterator.Root.(Utils.Node).Segment.Lexema

	iterator = iterator.Next
	parameters, errs = getParams(iterator, table)
	if _, ok := parameters[variable.Identifier]; ok {
		errs = append(errs, fmt.Errorf("el parametro %s ya se encuentra declarado", variable.Identifier))
		return
	}
	parameters[variable.Identifier] = variable
	table.Set(&variable, nil)
	return
}

func defFunc(tree *Utils.Tree, table *Table) []error {
	var localErrors []error
	fun := Function{}
	functionTable := Table{}

	iterator := tree.Children
	fun.ReturnType = iterator.Root.(Utils.Node).Segment.Lexema

	iterator = iterator.Next
	fun.Identifier = iterator.Root.(Utils.Node).Segment.Lexema

	iterator = iterator.Next.Next
	paramPointer := iterator

	iterator = iterator.Next.Next
	functionBlock := iterator

	functionTable.Stack = table.Stack
	functionTable.Stack = table.dumpStack()

	parameters, errs := getParams(paramPointer, &functionTable)
	fun.Parameters = parameters

	functionTable.Stack.Push(&fun)
	if len(errs) > 0 {
		localErrors = append(localErrors, errs...)
	}

	if table.Includes(fun.Identifier) {
		localErrors = append(localErrors, fmt.Errorf("la función %s ya se encuentra declarada", fun.Identifier))
	}

	table.Set(&fun, &functionTable)
	return append(localErrors, localDef(functionBlock.Children.Next, &functionTable, fun)...)

}
