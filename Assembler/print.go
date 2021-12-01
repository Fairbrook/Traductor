package Assembler

import "github.com/Fairbrook/analizador/Semantico"

func (t *Translator) tPrint(expression Semantico.Expression, table Semantico.Table, scope Semantico.Symbol) {
	functionName := expression.Parts[0].GetIdentifier()
	replace := specialFunctions[functionName]
	if expression.Parts[1].GetReturnType() == "float" {
		t.tExpression(*expression.Parts[1].(*Semantico.Expression), table, scope)
		t.stream.Write([]byte("\nfstp _feax"))
		t.stream.Write([]byte("\nprintf(\"%f\", _feax)"))
		return
	}
	if expression.Parts[1].GetReturnType() != "char*" {
		t.tExpression(*expression.Parts[1].(*Semantico.Expression), table, scope)
		t.stream.Write([]byte("\n" + replace + "eax)"))
		return
	}
	t.stream.Write([]byte("\n" + replace + expression.Parts[1].GetSegment().Lexema + ")"))
}
