package Assembler

import (
	"fmt"
	"math"
	"strconv"

	"github.com/Fairbrook/analizador/Semantico"
)

func (t *Translator) tCall(expression Semantico.Expression, table Semantico.Table, scope Semantico.Symbol) {
	//Funciones de print especiales porque soy flojo
	functionName := expression.Parts[0].GetIdentifier()
	if _, ok := specialFunctions[functionName]; ok {
		t.tPrint(expression, table, scope)
		return
	}

	t.stream.Write([]byte("\n; Llamada a la función " + functionName))
	for index := len(expression.Parts) - 1; index >= 1; index-- {
		t.tExpression(*expression.Parts[index].(*Semantico.Expression), table, scope)
		if expression.Parts[index].GetReturnType() == "float" {
			t.stream.Write([]byte("\nfstp _fhelper"))
			t.stream.Write([]byte("\npush _fhelper"))
			continue
		}
		t.stream.Write([]byte("\npush eax"))
	}
	t.stream.Write([]byte("\ncall " + functionName))
}

func tLiteral(literal Semantico.Literal, table Semantico.Table, scope Semantico.Symbol) []byte {
	res := ""
	if literal.GetReturnType() == "float" {
		real, _ := strconv.ParseFloat(literal.Value, 32)
		res = "\nmov _fhelper, " + fmt.Sprintf("%b", math.Float32bits(float32(real))) + "b"
		res += "\nfld _fhelper"
		return []byte(res)
	}
	return []byte("\nmov eax, " + literal.Value)
}

func (t *Translator) tTermino(expression Semantico.Expression, table Semantico.Table, scope Semantico.Symbol) {
	term := expression.Parts[0]
	if term.GetType() == "literal" {
		t.stream.Write(tLiteral(*term.(*Semantico.Literal), table, scope))
		return
	}
	id := expression.GetIdentifier()
	if _, local := table.Get(expression.GetIdentifier()); local {
		id = getSubScope(scope, &expression)
	}
	if expression.GetReturnType() == "float" {
		t.stream.Write([]byte("\nfld " + id))
		return
	}
	t.stream.Write([]byte("\nmov eax, " + id))
	return
}

func (t *Translator) tVariable(expression Semantico.Expression, table Semantico.Table, scope Semantico.Symbol) {
	if _, local := table.Get(expression.GetIdentifier()); local {
		t.stream.Write([]byte("\nmov eax, " + getSubScope(scope, &expression)))
		return
	}
	t.stream.Write([]byte("\nmov eax, " + expression.GetIdentifier()))
}

func (t *Translator) tReturn(expression Semantico.Expression, table Semantico.Table, scope Semantico.Symbol) {
	t.tExpression(*expression.Parts[0].(*Semantico.Expression), table, scope)
	if scope.GetIdentifier() == "main" {
		t.stream.Write([]byte("\n; Retorno al SO "))
		t.stream.Write([]byte("\ninvoke ExitProcess, eax"))
		return
	}
	t.stream.Write([]byte("\nret"))
}

func (t *Translator) tAsign(expression Semantico.Expression, table Semantico.Table, scope Semantico.Symbol) {
	t.tExpression(*expression.Parts[1].(*Semantico.Expression), table, scope)
	variable := expression.Parts[0].(*Semantico.Variable)
	identifier := variable.GetIdentifier()
	if _, local := table.Get(variable.GetIdentifier()); local {
		identifier = getSubScope(scope, variable)
	}
	if variable.GetReturnType() == "float" {
		t.stream.Write([]byte("\nfstp " + identifier))
		return
	}
	t.stream.Write([]byte("\nmov " + identifier + ", eax"))
}

func (t *Translator) tExpression(expression Semantico.Expression, table Semantico.Table, scope Semantico.Symbol) {
	subType := expression.Subtype
	switch subType {
	case "LlamadaFunc":
		t.tCall(expression, table, scope)
	case "OP":
		t.tOP(expression, table, scope)
	case "Termino":
		t.tTermino(expression, table, scope)
	case "return":
		t.tReturn(expression, table, scope)
	case "asign":
		t.tAsign(expression, table, scope)
	}
}
