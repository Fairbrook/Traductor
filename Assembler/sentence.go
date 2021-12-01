package Assembler

import (
	"github.com/Fairbrook/analizador/Semantico"
)

func (t *Translator) tSentence(symbol Semantico.Symbol, table Semantico.Table, scope Semantico.Symbol) {
	st := symbol.GetType()
	switch st {
	case "expression":
		{
			expression := *symbol.(*Semantico.Expression)
			if expression.Subtype == "LlamadaFunc" && expression.Parts[0].GetReturnType() == "float" {
				t.tCall(expression, table, scope)
				t.stream.Write([]byte("\nfstp _fhelper"))
				return
			}
			t.tExpression(expression, table, scope)
			return
		}
	case "conditional":
		t.tConditional(*symbol.(*Semantico.Conditional), table, scope)
	case "return":
		t.tReturn(*symbol.(*Semantico.Expression), table, scope)
	default:
		return
	}
}
