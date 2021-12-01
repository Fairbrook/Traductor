package Assembler

import "github.com/Fairbrook/analizador/Semantico"

func (t *Translator) tDefVariable(variable Semantico.Variable) {
	intialValue := "0"
	if variable.GetReturnType() == "float" {
		intialValue = "0.0"
	}
	t.stream.Write([]byte("\n" + variable.Identifier + " " + translateTypes[variable.GetReturnType()] + " " + intialValue))
}
