package Assembler

import (
	"errors"
	"os"
	"strconv"

	"github.com/Fairbrook/analizador/Semantico"
)

func (t *Translator) header() {
	t.stream.Write([]byte(".386\n"))
	t.stream.Write([]byte(".model flat, stdcall\n"))
	t.stream.Write([]byte("option casemap:none\n\n"))
	t.stream.Write([]byte("INCLUDE \\masm32\\include\\masm32rt.inc\n"))
	t.stream.Write([]byte("\n.data"))
	t.stream.Write([]byte("\n_feax real8 0.0"))
	t.stream.Write([]byte("\n_fhelper dword 0\n"))
}

func (t *Translator) footer() {
	t.stream.Write([]byte("\n\nend main"))
}

func sortSentences(sentences []Semantico.Symbol) (variables []Semantico.Symbol, functions []Semantico.Symbol) {
	variables = make([]Semantico.Symbol, 0)
	functions = make([]Semantico.Symbol, 0)
	for _, sentence := range sentences {
		sentenceType := sentence.GetType()
		if sentenceType == "function" {
			functions = append(functions, sentence)
			continue
		}
		variables = append(functions, sentence)
		continue
	}
	return
}

func (t *Translator) nextIfTag() (if_tag string, else_tag string, end_if string) {
	counter := strconv.FormatInt(t.tags_counter["if"], 10)
	if_tag = "_if_" + counter
	else_tag = "_else_" + counter
	end_if = "_endif_" + counter
	t.tags_counter["if"]++
	return
}

func (t *Translator) nextWhileTag() (while_tag string, end_while string) {
	counter := strconv.FormatInt(t.tags_counter["while"], 10)
	while_tag = "_while_" + counter
	end_while = "_endw_" + counter
	t.tags_counter["while"]++
	return
}

func (t *Translator) Translate(str string) (errs []error) {
	table, errs, sentences := Semantico.Analize(str)
	if len(errs) > 0 {
		return
	}
	file, e := os.Create(t.Filename)
	if e != nil {
		errs = []error{
			errors.New("error al crear el archivo ensamblador"),
		}
		return
	}
	t.stream = file
	t.header()
	variables, functions := sortSentences(sentences)
	t.tags_counter = map[string]int64{
		"if":    0,
		"while": 0,
	}
	for _, sentence := range variables {
		t.tDefVariable(*sentence.(*Semantico.Variable))
	}
	t.stream.Write([]byte("\n.code"))
	for _, sentence := range functions {
		t.tFunction(*sentence.(*Semantico.Function), table)
	}
	t.footer()
	file.Close()
	return
}
