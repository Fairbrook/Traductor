package Lexico

import (
	"errors"
	"fmt"
	"strings"
)

type Lexico struct {
	index     int
	prevIndex int
	Input     string
}

func (lex *Lexico) NextSegment() (segment Segment, err error) {
	if lex.index >= len(lex.Input) {
		err = errors.New("Index fuera de rango")
		return
	}
	segment, err = evaluate(lex.Input[lex.index:])
	lex.prevIndex = lex.index
	lex.index += segment.Index
	if err != nil {
		err = errors.New(err.Error() + fmt.Sprintf(" en la linea %d", strings.Count(lex.Input[0:lex.index], "\n")+1))
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
	lex.index = lex.prevIndex
}
