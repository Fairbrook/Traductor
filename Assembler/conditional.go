package Assembler

import (
	"github.com/Fairbrook/analizador/Semantico"
)

func (t *Translator) tWhile(con Semantico.Conditional, table Semantico.Table, scope Semantico.Symbol) {
	tag, end_tag := t.nextWhileTag()
	t.stream.Write([]byte("\n\n; While " + con.Expression.GetIdentifier()))
	t.stream.Write([]byte("\n" + tag + ":"))
	t.tExpression(con.Expression, table, scope)
	t.stream.Write([]byte("\njz " + end_tag))
	t.stream.Write([]byte("\n; then:"))
	for _, sentence := range con.Sentences {
		t.tSentence(sentence, table, scope)
	}
	t.stream.Write([]byte("\njmp " + tag))
	t.stream.Write([]byte("\n" + end_tag + ":"))
	t.stream.Write([]byte("\n; End While\n"))
}

func (t *Translator) tIf(con Semantico.Conditional, table Semantico.Table, scope Semantico.Symbol) {
	_, else_tag, end_tag := t.nextIfTag()
	t.stream.Write([]byte("\n\n; if " + con.Expression.GetIdentifier()))
	t.tExpression(con.Expression, table, scope)
	nextTag := end_tag
	if con.HasAlt() {
		nextTag = else_tag
	}
	t.stream.Write([]byte("\njz " + nextTag))
	t.stream.Write([]byte("\n; then:"))
	for _, sentence := range con.Sentences {
		t.tSentence(sentence, table, scope)
	}
	if con.HasAlt() {
		t.stream.Write([]byte("\njmp " + end_tag))
		t.stream.Write([]byte("\n" + else_tag + ":"))
	}
	if con.HasAlt() {
		for _, sentence := range con.AltSentences {
			t.tSentence(sentence, table, scope)
		}
	}
	t.stream.Write([]byte("\n" + end_tag + ":"))
	t.stream.Write([]byte("\n; End if\n"))
}

func (t *Translator) tConditional(con Semantico.Conditional, table Semantico.Table, scope Semantico.Symbol) {
	switch con.Subtype {
	case "while":
		t.tWhile(con, table, scope)
	case "if":
		t.tIf(con, table, scope)
	}
}
