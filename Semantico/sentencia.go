package Semantico

import (
	"github.com/Fairbrook/analizador/Utils"
)

func procTermino(tree *Utils.Tree, table *Table) ([]error, Expression) {
	errors := make([]error, 0)
	iterator := tree.Children
	term := iterator.Root.(Utils.Node).Segment
	var symbol Symbol
	if term.Lexema == "LlamadaFunc" {
		return procCall(iterator, table)
	}
	switch term.Type {
	case 0:
		{
			termData, _ := table.Get(term.Lexema)
			if termData != nil {
				if termData.GetType() != VariableType {
					errors = append(errors, &Utils.SegmentError{
						Segment: term,
						Message: Utils.NoVarMsg,
					})
					break
				}
				symbol = termData.(*Variable)
			} else {
				errors = append(errors, &Utils.SegmentError{
					Segment: term,
					Message: Utils.NoDefMsg,
				})
			}
			break
		}
	case 1:
		symbol = &Literal{
			Value:      term.Lexema,
			ReturnType: "int",
			Segment:    term,
		}
	case 2:
		symbol = &Literal{
			Value:      term.Lexema,
			ReturnType: "float",
			Segment:    term,
		}
	case 3:
		symbol = &Literal{
			Value:      term.Lexema,
			ReturnType: "char*",
			Segment:    term,
		}
	}

	return errors, Expression{
		ReturnType: symbol.GetReturnType(),
		Subtype:    "Termino",
		Parts:      []Symbol{symbol},
	}
}

func procExpresion(tree *Utils.Tree, table *Table) ([]error, Expression) {
	errors := make([]error, 0)
	var retType string
	if tree.Children == nil {
		return errors, Expression{}
	}
	iterator := tree.Children
	lexema := iterator.Root.(Utils.Node).Segment.Lexema
	if lexema == "Termino" {
		return procTermino(iterator, table)
	}

	if lexema == "Expresion" {
		expression := Expression{
			Subtype: "OP",
		}
		e, ex := procExpresion(iterator, table)
		errors = append(errors, e...)
		expression.Parts = []Symbol{&ex}

		iterator = iterator.Next
		operatorSeg := iterator.Root.(Utils.Node).Segment
		operator := OpByType[operatorSeg.Type]
		expression.Parts = append(expression.Parts, &OperatorSymbol{
			Operator: operator,
			Segment:  operatorSeg,
		})

		retType = ex.GetReturnType()
		iterator = iterator.Next
		if (!operator.isBinary && iterator != nil) ||
			(operator.isBinary && iterator == nil) ||
			!operator.acceptType(ex.GetReturnType()) {
			errors = append(errors, &Utils.SegmentError{
				Segment: operatorSeg,
				Message: Utils.WrongOpMsg,
			})
			return errors, expression
		}

		e2, ex2 := procExpresion(iterator, table)
		errors = append(errors, e2...)
		expression.Parts = append(expression.Parts, &ex2)

		if ex2.GetReturnType() != "" && retType != ex2.GetReturnType() {
			errors = append(errors, &Utils.SegmentError{
				Segment: operatorSeg,
				Message: Utils.WrongOpMsg,
			})
		}
		if operator.returnType != "same" {
			retType = operator.returnType
		}
		expression.ReturnType = retType
		return errors, expression
	}
	return errors, Expression{}
}

func getArgs(tree *Utils.Tree, table *Table) ([]error, []Expression) {
	errors := make([]error, 0)
	values := make([]Expression, 0)
	if tree == nil || tree.Children == nil {
		return errors, values
	}
	iterator := tree.Children
	if iterator.Root.(Utils.Node).Segment.Lexema == "," {
		iterator = iterator.Next
	}
	var err []error
	var ex Expression
	err, ex = procExpresion(iterator, table)
	errors = append(errors, err...)
	values = append(values, ex)

	err, retArray := getArgs(iterator.Next, table)
	errors = append(errors, err...)
	values = append(values, retArray...)
	return errors, values
}

func procCall(tree *Utils.Tree, table *Table) ([]error, Expression) {
	errors := make([]error, 0)
	call := tree.Children.Root.(Utils.Node).Segment
	id := call.Lexema
	arguPointer := tree.Children.Next.Next
	funSym, _ := table.Get(id)
	expression := Expression{}
	if funSym == nil {
		errors = append(errors, &Utils.SegmentError{
			Segment: call,
			Message: Utils.NoDefMsg,
		})
		return errors, expression
	}
	fun := funSym.(*Function)
	expression.Parts = []Symbol{fun}
	expression.ReturnType = fun.ReturnType
	expression.Subtype = "LlamadaFunc"

	err, arguments := getArgs(arguPointer, table)
	errors = append(errors, err...)
	if len(arguments) < len(fun.Parameters) {
		errors = append(errors, &Utils.SegmentError{
			Segment: call,
			Message: Utils.ArgsMsg,
		})
		return errors, expression
	}
	if len(arguments) > len(fun.Parameters) {
		errors = append(errors, &Utils.SegmentError{
			Segment: call,
			Message: Utils.ArgsMsg,
		})
		return errors, expression
	}

	i := 0
	for _, param := range fun.ParamList {
		if param.Type != arguments[i].GetReturnType() {
			errors = append(errors, &Utils.SegmentError{
				Segment: call,
				Message: Utils.ArgsMsg,
			})
			break
		}
		expression.Parts = append(expression.Parts, &arguments[i])
		i++
	}

	return errors, expression
}

func procWhile(tree *Utils.Tree, table *Table, function Function) ([]error, Conditional) {
	errors := make([]error, 0)
	conditional := Conditional{
		Subtype: "while",
	}
	iterator := tree.Children.Next.Next
	e, ex := procExpresion(iterator, table)
	conditional.Expression = ex
	errors = append(errors, e...)
	if ex.GetReturnType() != "bool" {
		errors = append(errors, &Utils.SegmentError{
			Segment: iterator.Root.(Utils.Node).Segment,
			Message: Utils.NoBoolMsg,
		})
	}
	iterator = iterator.Next.Next
	sentences := iterator.Children.Next

	blockTable := Table{}
	blockTable.Stack = table.dumpStack()

	e, conditional.Sentences = procSentencias(sentences, &blockTable, function)
	errors = append(errors, e...)

	return errors, conditional
}
func procIf(tree *Utils.Tree, table *Table, function Function) ([]error, Conditional) {
	errors := make([]error, 0)
	conditional := Conditional{
		Subtype: "if",
	}

	iterator := tree.Children.Next.Next
	e, ex := procExpresion(iterator, table)
	conditional.Expression = ex

	errors = append(errors, e...)
	if ex.GetReturnType() != "bool" {
		errors = append(errors, &Utils.SegmentError{
			Segment: iterator.Root.(Utils.Node).Segment,
			Message: Utils.NoBoolMsg,
		})
	}
	iterator = iterator.Next.Next
	sentences := iterator.Children.Children.Next

	blockTable := Table{}
	blockTable.Stack = table.dumpStack()

	e, conditional.Sentences = procSentencias(sentences, &blockTable, function)
	errors = append(errors, e...)

	iterator = iterator.Next
	if iterator.Children == nil {
		return errors, conditional
	}
	sentences = iterator.Children.Next.Children.Children.Next

	elseTable := Table{}
	elseTable.Stack = table.dumpStack()
	e, conditional.AltSentences = procSentencias(sentences, &blockTable, function)
	errors = append(errors, e...)

	return errors, conditional
}

func procSentencias(tree *Utils.Tree, table *Table, function Function) ([]error, []Symbol) {
	errors := make([]error, 0)
	symbols := make([]Symbol, 0)
	if tree.Children == nil {
		return errors, symbols
	}
	err, sym := procSentencia(tree.Children, table, function)
	errors = append(errors, err...)
	symbols = append(symbols, sym...)
	err, sym = procSentencias(tree.Children.Next, table, function)
	errors = append(errors, err...)
	symbols = append(symbols, sym...)

	return errors, symbols
}

func procReturn(tree *Utils.Tree, table *Table, function Function) ([]error, []Symbol) {
	errors := make([]error, 0)
	err, ex := procExpresion(tree.Next.Children, table)
	errors = append(errors, err...)
	if ex.GetReturnType() != function.ReturnType {
		seg := tree.Next.Children.Root.(Utils.Node).Segment
		errors = append(errors, &Utils.SegmentError{
			Segment: Utils.Segment{
				Lexema: function.GetIdentifier(),
				Line:   seg.Line,
				Index:  seg.Index,
			},
			Message: Utils.RetMsg,
		})
	}

	return errors, []Symbol{
		&Expression{
			ReturnType: function.GetReturnType(),
			Subtype:    "return",
			Parts:      []Symbol{&ex},
		},
	}
}

func procSentencia(tree *Utils.Tree, table *Table, function Function) ([]error, []Symbol) {
	errors := make([]error, 0)
	iterator := tree.Children
	lexema := iterator.Root.(Utils.Node).Segment.Lexema
	switch lexema {
	case "LlamadaFunc":
		{
			err, ex := procCall(iterator, table)
			errors = append(errors, err...)
			return errors, []Symbol{&ex}
		}
	case "if":
		{
			err, con := procIf(tree, table, function)
			errors = append(errors, err...)
			return errors, []Symbol{&con}
		}
	case "while":
		{
			err, con := procWhile(tree, table, function)
			errors = append(errors, err...)
			return errors, []Symbol{&con}
		}
	case "return":
		{
			err, con := procReturn(iterator, table, function)
			errors = append(errors, err...)
			return errors, con
		}
	default:
		{
			variable, _ := table.Get(lexema)
			err, ex := procExpresion(iterator.Next.Next, table)
			errors = append(errors, err...)
			if variable == nil {
				errors = append(errors, &Utils.SegmentError{
					Segment: iterator.Root.(Utils.Node).Segment,
					Message: Utils.NoDefMsg,
				})
				return errors, []Symbol{}
			}
			if ex.GetReturnType() != "" && variable.(*Variable).Type != ex.GetReturnType() {
				errors = append(errors, &Utils.SegmentError{
					Segment: iterator.Root.(Utils.Node).Segment,
					Message: Utils.WrongTypeMsg,
				})
			}
			return errors, []Symbol{
				&Expression{
					ReturnType: ex.GetReturnType(),
					Subtype:    "asign",
					Parts: []Symbol{
						variable,
						&ex,
					},
				},
			}
		}
	}
}
