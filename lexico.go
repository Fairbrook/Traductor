package main

import (
	"errors"
	"fmt"
	"strings"
)

var reserved = map[string]byte{
	"return": 1,
	"else":   2,
	"if":     3,
	"while":  4,
}

var tipo = map[string]byte{
	"int":   1,
	"float": 2,
	"void":  3,
}

var NumCodes = map[string]uint8{
	"Identificador": 0,
	"Entero":        1,
	"Decimal":       2,
	"Cadena":        3,
	"Tipo":          4,
	"OpSuma":        5,
	"OpMul":         6,
	"OpRelac":       7,
	"OpOr":          8,
	"OpAnd":         9,
	"OpNot":         10,
	"OpIgualdad":    11,
	";":             12,
	",":             13,
	"(":             14,
	")":             15,
	"{":             16,
	"}":             17,
	"=":             18,
	"if":            19,
	"while":         20,
	"return":        21,
	"else":          22,
	"$":             23,
}

func getSpecialType(in Segment) (out Segment) {
	out.Lexema = in.Lexema
	out.StateName = in.StateName
	if in.State == 1 {
		if _, ok := reserved[in.Lexema]; ok {
			out.StateName = out.Lexema
		}
		if _, ok := tipo[in.Lexema]; ok {
			out.StateName = "Tipo"
		}
	}
	if in.State == 11 {
		out.StateName = out.Lexema
	}
	out.State = NumCodes[out.StateName]
	return
}

func processString(str string) (res []Segment, err error) {
	index := 0
	var segment Segment
	input := str + "$"
	for index < len(input) {
		segment, err = evaluate(input[index:])
		index += segment.Index
		if err != nil {
			err = errors.New(err.Error() + fmt.Sprintf(" en la linea %d", strings.Count(input[0:index], "\n")+1))
			return
		}
		res = append(res, getSpecialType(segment))
	}
	return
}
