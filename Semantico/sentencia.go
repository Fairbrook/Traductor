package Semantico

import (
	"fmt"

	"github.com/Fairbrook/analizador/Utils"
)

func procTermino(tree *Utils.Tree, table *Table) ([]error, string) {
	errors := make([]error, 0)
	var retType string
	iterator := tree.Children
	term := iterator.Root.(Utils.Node).Segment
	if term.Lexema == "LlamadaFunc" {
		return procCall(iterator, table)
	}
	switch term.Type {
	case 0:
		{
			termData := table.Get(term.Lexema)
			if termData != nil {
				retType = termData.(*Variable).Type
			} else {
				errors = append(errors, fmt.Errorf("el identificador %s no esta definido", term.Lexema))
			}
			break
		}
	case 1:
		retType = "int"
	case 2:
		retType = "float"
	case 3:
		retType = "char*"
	}

	return errors, retType
}

func procExpresion(tree *Utils.Tree, table *Table) ([]error, string) {
	errors := make([]error, 0)
	var retType string
	if tree.Children == nil {
		return errors, retType
	}
	iterator := tree.Children
	lexema := iterator.Root.(Utils.Node).Segment.Lexema
	if lexema == "Termino" {
		return procTermino(iterator, table)
	}

	if lexema == "Expresion" {
		e, t := procExpresion(iterator, table)
		errors = append(errors, e...)
		iterator = iterator.Next
		operatorSeg := iterator.Root.(Utils.Node).Segment
		operator := OpByType[operatorSeg.Type]
		retType = t
		iterator = iterator.Next
		if (!operator.isBinary && iterator != nil) ||
			(operator.isBinary && iterator == nil) ||
			!operator.acceptType(t) {
			errors = append(errors, fmt.Errorf("expresion invalida"))
			return errors, retType
		}
		e, t = procExpresion(iterator, table)
		errors = append(errors, e...)
		if t != "" && retType != t {
			errors = append(errors, fmt.Errorf("los tipos en la expresion no coinciden"))
		}
		if operator.returnType != "same" {
			retType = operator.returnType
		}
		return errors, retType
	}
	return errors, retType
}

func getArgs(tree *Utils.Tree, table *Table) ([]error, []string) {
	errors := make([]error, 0)
	values := make([]string, 0)
	if tree.Children == nil {
		return errors, values
	}
	err, retArray := getArgs(tree.Children, table)
	errors = append(errors, err...)
	values = append(values, retArray...)
	var retType string
	err, retType = procExpresion(tree.Children.Next, table)
	errors = append(errors, err...)
	values = append(values, retType)
	return errors, values
}

func procCall(tree *Utils.Tree, table *Table) ([]error, string) {
	errors := make([]error, 0)
	id := tree.Children.Root.(Utils.Node).Segment.Lexema
	arguPointer := tree.Children.Next.Next
	funSym := table.Get(id)
	if funSym == nil {
		errors = append(errors, fmt.Errorf("no se encontro una definicion de la funcion %s", id))
		return errors, ""
	}
	fun := funSym.(*Function)
	err, arguments := getArgs(arguPointer, table)
	i := 0
	errors = append(errors, err...)
	for _, param := range fun.Parameters {
		if param.Type != arguments[i] {
			errors = append(errors, fmt.Errorf("la llamada a la funcion %s no recuerda co la declaracion", fun.Identifier))
			break
		}
		i++
	}
	return errors, fun.ReturnType
}

func procWhile(tree *Utils.Tree, table *Table, function Function) []error {
	errors := make([]error, 0)
	iterator := tree.Children.Next.Next
	e, retType := procExpresion(iterator, table)
	errors = append(errors, e...)
	if retType != "bool" {
		errors = append(errors, fmt.Errorf("expresion no booleana en if"))
	}
	iterator = iterator.Next
	sentences := iterator.Children.Next

	blockTable := Table{}
	blockTable.Stack = table.dumpStack()

	e = procSentencias(sentences, &blockTable, function)
	errors = append(errors, e...)

	return errors
}
func procIf(tree *Utils.Tree, table *Table, function Function) []error {
	errors := make([]error, 0)
	iterator := tree.Children.Next.Next
	e, retType := procExpresion(iterator, table)
	errors = append(errors, e...)
	if retType != "bool" {
		errors = append(errors, fmt.Errorf("expresion no booleana en if"))
	}
	iterator = iterator.Next.Next
	sentences := iterator.Children.Children.Next

	blockTable := Table{}
	blockTable.Stack = table.dumpStack()

	e = procSentencias(sentences, &blockTable, function)
	errors = append(errors, e...)

	iterator = iterator.Next
	if iterator.Children == nil {
		return errors
	}
	sentences = iterator.Children.Next.Children.Children.Next

	elseTable := Table{}
	elseTable.Stack = table.dumpStack()
	e = procSentencias(sentences, &blockTable, function)
	errors = append(errors, e...)

	return errors
}

func procSentencias(tree *Utils.Tree, table *Table, function Function) []error {
	errors := make([]error, 0)
	if tree.Children == nil {
		return errors
	}
	err := procSentencia(tree.Children, table, function)
	errors = append(errors, err...)
	err = procSentencias(tree.Children.Next, table, function)
	errors = append(errors, err...)
	return errors
}

func procSentencia(tree *Utils.Tree, table *Table, function Function) []error {
	errors := make([]error, 0)
	iterator := tree.Children
	lexema := iterator.Root.(Utils.Node).Segment.Lexema
	switch lexema {
	case "LlamadaFunc":
		{
			err, _ := procCall(iterator, table)
			errors = append(errors, err...)
			return errors
		}
	case "if":
		{
			err := procIf(tree, table, function)
			errors = append(errors, err...)
			return errors
		}
	case "while":
		{
			err := procWhile(tree, table, function)
			errors = append(errors, err...)
			return errors
		}
	case "return":
		{
			err, retType := procExpresion(iterator.Next.Next, table)
			errors = append(errors, err...)
			if retType != function.ReturnType {
				errors = append(errors, fmt.Errorf("el valor de retorno no corresponde con el valor definido de la funcion %s", function.getIdentifier()))
			}
			return errors
		}
	default:
		{
			variable := table.Get(lexema)
			err, retType := procExpresion(iterator.Next.Next, table)
			errors = append(errors, err...)
			if variable == nil {
				errors = append(errors, fmt.Errorf("la variable %s no esta definida", variable))
				return errors
			}
			if retType != "" && variable.(*Variable).Type != retType {
				errors = append(errors, fmt.Errorf("la variable %s no es de tipo %s", variable.getIdentifier(), retType))
			}
			return errors
		}
	}
}
