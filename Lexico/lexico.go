package Lexico

import (
	"errors"

	"github.com/Fairbrook/analizador/Utils"
)

type Lexico struct {
	index       int
	prevIndex   int
	currentLine int
	prevLine    int
	Input       string
}

func (lex *Lexico) NextSegment() (segment Utils.Segment, err error) {
	if lex.index >= len(lex.Input) {
		err = errors.New("Index fuera de rango")
		return
	}
	segment, err = evaluate(lex.Input[lex.index:], lex.currentLine)
	lex.prevIndex = lex.index
	lex.index += segment.Index
	lex.prevLine = lex.currentLine
	lex.currentLine = segment.Line
	segment.Line = segment.Line + 1
	if err != nil {
		return
	}
	segment = getSpecialType(segment)
	return
}

func (lex *Lexico) EOF() bool {
	return lex.index >= len(lex.Input)
}

func (lex *Lexico) GetLast() string {
	return lex.Input[lex.index:]
}

func (lex *Lexico) GoBack() {
	lex.currentLine = lex.prevLine
	lex.index = lex.prevIndex
}
