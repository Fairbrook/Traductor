package Semantico

import (
	"github.com/Fairbrook/analizador/Utils"
)

func localDef(tree *Utils.Tree, table *Table, function Function) ([]error, []Symbol) {
	if tree.Children == nil {
		return make([]error, 0), make([]Symbol, 0)
	}
	defType := tree.Root.(Utils.Node).Segment.Lexema
	switch defType {
	case "DefLocales":
		{
			errors, simbols := localDef(tree.Children, table, function)
			nextErrors, nextSimbols := localDef(tree.Children.Next, table, function)
			errors = append(errors, nextErrors...)
			simbols = append(simbols, nextSimbols...)
			return errors, simbols
		}
	case "DefLocal":
		return localDef(tree.Children, table, function)
	case "DefVar":
		return defVar(tree, table, "")
	case "Sentencia":
		return procSentencia(tree, table, function)
	default:
		return make([]error, 0), make([]Symbol, 0)

	}
}

func getParams(params *Utils.Tree, table *Table) (parameters map[string]Variable, list []Variable, errs []error) {
	list = make([]Variable, 0)
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
	variable.Segment = iterator.Root.(Utils.Node).Segment
	variable.Identifier = variable.Segment.Lexema
	list = []Variable{variable}

	iterator = iterator.Next
	var rest []Variable
	parameters, rest, errs = getParams(iterator, table)
	list = append(list, rest...)
	if _, ok := parameters[variable.Identifier]; ok {
		errs = append(errs, &Utils.SegmentError{
			Segment: table.Parent.GetSegment(),
			Message: Utils.DeclaredMsg,
		})
		return
	}
	parameters[variable.Identifier] = variable
	table.Set(&variable, nil)
	return
}

func defFunc(tree *Utils.Tree, table *Table) ([]error, []Symbol) {
	var localErrors []error
	fun := Function{}
	functionTable := Table{}

	iterator := tree.Children
	fun.ReturnType = iterator.Root.(Utils.Node).Segment.Lexema

	iterator = iterator.Next
	fun.Segment = iterator.Root.(Utils.Node).Segment
	fun.Identifier = fun.Segment.Lexema

	iterator = iterator.Next.Next
	paramPointer := iterator

	iterator = iterator.Next.Next
	functionBlock := iterator

	functionTable.Stack = table.Stack
	functionTable.Stack = table.dumpStack()

	functionTable.Parent = &fun

	parameters, list, errs := getParams(paramPointer, &functionTable)
	fun.Parameters = parameters
	fun.ParamList = list

	if len(errs) > 0 {
		localErrors = append(localErrors, errs...)
	}

	if table.Includes(fun.Identifier, false) {
		localErrors = append(localErrors, &Utils.SegmentError{
			Segment: fun.Segment,
			Message: Utils.DeclaredMsg,
		})
	}

	table.Set(&fun, &functionTable)
	errors, simbols := localDef(functionBlock.Children.Next, &functionTable, fun)
	errors = append(localErrors, errors...)
	fun.Sentences = simbols
	if !fun.HasReturn() && fun.ReturnType != "void" {
		errors = append(errors, &Utils.SegmentError{
			Segment: fun.Segment,
			Message: Utils.NoReturnMsg,
		})
	}
	simbols = []Symbol{&fun}
	return errors, simbols
}
