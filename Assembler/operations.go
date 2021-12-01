package Assembler

import "github.com/Fairbrook/analizador/Semantico"

var flagsMask = "0100011100000000B"
var equalsMask = "0100000000000000B"
var greaterMask = "0000000000000000B"
var lessMask = "0000000100000000B"
var tCompare = "\nfcom\nfstsw ax\nand eax, " + flagsMask
var tPop2 = "\n;limpiar fpu stack\nfstp _fhelper\nfstp _fhelper"
var aritm = map[string]string{
	"+": "\nfadd",
	"-": "\nfsub",
	"*": "\nfimul",
	"/": "\nfidiv",
}

func (t *Translator) tOpRealAritm(expression Semantico.Expression, table Semantico.Table, scope Semantico.Symbol) {
	t.tExpression(*expression.Parts[0].(*Semantico.Expression), table, scope)
	t.tExpression(*expression.Parts[2].(*Semantico.Expression), table, scope)
	opLexema := expression.Parts[1].(*Semantico.OperatorSymbol).Segment.Lexema
	op := aritm[opLexema]
	t.stream.Write([]byte(op))
}

func (t *Translator) tOpReal(expression Semantico.Expression, table Semantico.Table, scope Semantico.Symbol) {
	op := expression.Parts[1].(*Semantico.OperatorSymbol).Segment.Lexema
	if _, ok := aritm[op]; ok {
		t.tOpRealAritm(expression, table, scope)
		return
	}
	t.tExpression(*expression.Parts[2].(*Semantico.Expression), table, scope)
	t.tExpression(*expression.Parts[0].(*Semantico.Expression), table, scope)
	operation := tCompare
	switch op {
	case "==":
		operation += "\ncmp eax, " + equalsMask
		operation += "\nsete al"
	case "!=":
		operation += "\ncmp eax, " + equalsMask
		operation += "\nsetne al"
	case ">":
		operation += "\ncmp eax, " + greaterMask
		operation += "\nsete al"
	case "<":
		operation += "\ncmp eax, " + lessMask
		operation += "\nsete al"
	case "<=":
		operation += "\nmov ebx, eax"
		operation += "\ncmp eax, " + equalsMask
		operation += "\nsete al"
		operation += "\ncmp ebx, " + lessMask
		operation += "\nsete bl"
		operation += "\nor al, bl"
	case ">=":
		operation += "\nmov ebx, eax"
		operation += "\ncmp eax, " + equalsMask
		operation += "\nsete al"
		operation += "\ncmp ebx, " + greaterMask
		operation += "\nsete bl"
		operation += "\nor al, bl"
	}
	operation += tPop2
	operation += "\nand eax, 000000FFh"
	t.stream.Write([]byte(operation))
}

func (t *Translator) tOP(expression Semantico.Expression, table Semantico.Table, scope Semantico.Symbol) {
	if expression.GetReturnType() == "float" || expression.Parts[0].GetReturnType() == "float" {
		t.tOpReal(expression, table, scope)
		return
	}
	t.tExpression(*expression.Parts[2].(*Semantico.Expression), table, scope)
	t.stream.Write([]byte("\nmov ebx, eax"))
	t.tExpression(*expression.Parts[0].(*Semantico.Expression), table, scope)
	op := expression.Parts[1].(*Semantico.OperatorSymbol).Segment.Lexema
	operation := ""
	switch op {
	case "-":
		operation = "\nsub eax, ebx"
	case "+":
		operation = "\nadd eax, ebx"
	case "/":
		operation = "\nidiv ebx"
	case "*":
		operation = "\nimul ebx"

	case "&&":
		operation = "\nand eax, ebx"

	case "||":
		operation = "\nor ax, ex"

	case "==":
		operation = "\nxor eax, ebx"
		operation += "\nsetz al"
		operation += "\nand eax, 000000FFh"

	case "!=":
		operation = "\nxor eax, ebx"
		operation += "\nsetnz al"
		operation += "\nand eax, 000000FFh"

	case ">":
		operation += "\nsub eax, ebx"
		operation += "\nsetg al"
		operation += "\nand eax, 000000FFh"

	case "<":
		operation += "\nsub eax, ebx"
		operation += "\nsetl al"
		operation += "\nand eax, 000000FFh"

	case "<=":
		operation += "\nsub eax, ebx"
		operation += "\nsetle al"
		operation += "\nand eax, 000000FFh"

	case ">=":
		operation += "\nsub eax, ebx"
		operation += "\nsetge al"
		operation += "\nand eax, 000000FFh"
	}
	t.stream.Write([]byte(operation))
}
