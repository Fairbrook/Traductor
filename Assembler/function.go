package Assembler

import (
	"github.com/Fairbrook/analizador/Semantico"
)

func (t *Translator) tFuncDecl(function Semantico.Function) {
	params := ""
	t.stream.Write([]byte("\n\n;---------- Funcion " + function.Identifier + " -----------\n"))
	for _, variable := range function.ParamList {
		if params != "" {
			params += ","
		}
		params += getSubScope(&function, &variable) + ":" + translateTypes[variable.GetReturnType()]
	}
	header := function.GetIdentifier() + " proc " + params
	t.stream.Write([]byte(header))
}

func (t *Translator) tFuncLocals(function Semantico.Function, subTable *Semantico.Table) {
	if subTable != nil {
		t.stream.Write([]byte("\n;locals"))
		for _, variable := range subTable.DumpTable() {
			if _, ok := function.Parameters[variable.GetIdentifier()]; !ok {
				t.stream.Write([]byte("\nlocal " + getSubScope(&function, variable) + ":" + translateTypes[variable.GetReturnType()]))
			}
		}
		t.stream.Write([]byte("\n;fin locals\n"))
	}
}

func (t *Translator) tFooter(function Semantico.Function) {
	if !function.HasReturn() {
		t.stream.Write([]byte("\nret"))
	}
	t.stream.Write([]byte("\n" + function.GetIdentifier() + " endp"))
}

func (t *Translator) tFunction(function Semantico.Function, table Semantico.Table) {
	t.tFuncDecl(function)
	subTable := table.GetSubTable(function.GetIdentifier())
	if subTable != nil {
		t.tFuncLocals(function, subTable)
	}
	if function.GetIdentifier() == "main" {
		t.stream.Write([]byte("\nfinit"))
	}
	for _, sentence := range function.Sentences {
		t.tSentence(sentence, *subTable, &function)
	}
	t.tFooter(function)
}
