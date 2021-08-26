package main

import (
	"errors"
	"fmt"
)

const noState uint8 = 255

var simbols = map[byte]byte{
	';': 1,
	',': 2,
	'(': 3,
	')': 4,
	'{': 5,
	'}': 6,
}

func isAlpha(character byte) bool {
	if character >= 'A' && character <= 'Z' {
		return true
	}
	if character >= 'a' && character <= 'z' {
		return true
	}
	return false
}

func isDigit(character byte) bool {
	if character >= '0' && character <= '9' {
		return true
	}
	return false
}

func isSpace(character byte) bool {
	if character == ' ' || character == '\n' || character == '\t' {
		return true
	}
	return false
}

func isValidString(character byte) bool {
	if character == '\n' || character == '"' {
		return false
	}
	return true
}

func isSimbol(character byte) bool {
	if _, ok := simbols[character]; ok {
		return true
	}
	return false
}

var functions = map[string]func(byte) bool{
	"isAlpha":       isAlpha,
	"isDigit":       isDigit,
	"isValidString": isValidString,
	"isSimbol":      isSimbol,
}

var states = map[uint8]map[string]uint8{
	0: {
		"isAlpha":  1,
		"isDigit":  2,
		"$":        3,
		"+":        4,
		"-":        4,
		"*":        5,
		"/":        5,
		"<":        6,
		">":        6,
		"|":        7,
		"&":        8,
		"!":        9,
		"=":        10,
		"isSimbol": 11,
		"\"":       12,
	},
	1: {
		"isAlpha": 1,
		"isDigit": 1,
	},
	2: {
		"isDigit": 2,
		".":       13,
	},
	3: {},
	4: {},
	5: {},
	6: {
		"=": 14,
	},
	7: {
		"|": 15,
	},
	8: {
		"&": 16,
	},
	9: {
		"=": 17,
	},
	10: {
		"=": 17,
	},
	11: {},
	12: {
		"isValidString": 12,
		"\"":            18,
	},
	13: {
		"isDigit": 19,
	},
	14: {},
	15: {},
	16: {},
	17: {},
	18: {},
	19: {
		"isDigit": 19,
	},
}

var terminals = map[uint8]string{
	1: "Identificador",
	2: "Entero",
	3: "$",
	4: "OpSuma",
	5: "OpMul",

	6:  "OpRelac",
	14: "OpRelac",

	9:  "OpNot",
	10: "=",
	11: "Simbolo",
	15: "OpOr",
	16: "OpAnd",
	17: "OpIgualdad",

	18: "Cadena",
	19: "Decimal",
}

type Segment struct {
	Lexema    string `json:"lexema"`
	State     uint8  `json:"state"`
	StateName string `json:"state_name"`
	Index     int    `json:"index"`
}

func evaluate(str string) (segment Segment, err error) {
	segment.Lexema = ""
	segment.StateName = ""
	segment.State = 0
	start := 0
	err = nil

	for i := 0; isSpace(str[i]) && len(str)-1 > i; i++ {
		start++
	}

	var nextState uint8
	for index := start; len(str) > index; index++ {
		nextState = noState

		for key, element := range states[segment.State] {
			if value, ok := functions[key]; ok {
				if value(str[index]) {
					nextState = element
					break
				}
				continue
			}

			if key[0] == str[index] && len(key) == 1 {
				nextState = element
				break
			}
		}

		if nextState == noState {
			segment.Index = index
			if name, ok := terminals[segment.State]; ok {
				segment.StateName = name
				return
			}
			err = errors.New(fmt.Sprintf("Caracter inesperado '%c'", str[index]))
			return
		}
		segment.State = nextState
		segment.Lexema += string(str[index])
	}

	segment.Index = len(str)
	if name, ok := terminals[segment.State]; ok {
		segment.StateName = name
		return
	}
	err = errors.New(fmt.Sprintf("Final de cadena inesperado al interpretar %s", str))

	return
}