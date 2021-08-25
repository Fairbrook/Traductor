package main

import (
	"errors"
	"fmt"
)

const noState uint8 = 255

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
	if character == ' ' || character == '\n' {
		return true
	}
	return false
}

var functions = map[string]func(byte) bool{
	"isAlpha": isAlpha,
	"isDigit": isDigit,
}

var states = map[uint8]map[string]uint8{
	0: {
		"isAlpha": 1,
		"isDigit": 2,
		"$":       3,
		// "isSimbol": 4,
		// "\"": 5,
	},
	1: {
		"isAlpha": 1,
		"isDigit": 1,
	},
	2: {
		"isDigit": 2,
		// ".":       3,
	},
	// 3: {
	// 	"isDigit": 3,
	// },
	// 4: {
	// 	"+": 6,
	// 	"-": 6,
	// 	"*": 7,
	// 	"/": 7,
	// 	"<": 8,
	// 	">": 8,
	// 	"=": 9,
	// 	"!": 10,
	// 	"|": 11,
	// 	"&": 12,
	// },
}

var terminals = map[uint8]string{
	1: "identificador",
	2: "entero",
	3: "$",
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
			}
			if key[0] == str[index] {
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
			err = errors.New(fmt.Sprintf("Caracter inesperado %c", str[index]))
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
