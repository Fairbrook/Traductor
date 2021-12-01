package Utils

import "fmt"

type SegmentError struct {
	Segment Segment
	Message string
}

var CaracterMsg string = "Caracter inesperado '%s' en la linea: %d"
var CadenaMsg string = "Cadena inesperada '%s' en la linea: %d"
var DeclaredMsg string = "El símbolo '%s' ya está declarado linea: %d"
var NoVarMsg string = "El símbolo '%s' no es una variable linea: %d"
var NoDefMsg string = "El símbolo '%s' no está definido linea: %d"
var ArgsMsg string = "Los argumentos no coinciden con la declaración de la función '%s' linea: %d"
var NoBoolMsg string = "Expresión no booleana en la condición de %s linea: %d"
var RetMsg string = "El valor de retorno no coincide con la declaración de '%s' linea: %d"
var WrongTypeMsg string = "El valor de la expresion '%s' no coincide linea: %d"
var WrongOpMsg string = "La expresión '%s' es inválida linea: %d"
var NoReturnMsg string = "La funcion '%s' debe tener un valor de retorno linea: %d"

func (se *SegmentError) Error() string {
	lexema := se.Segment.Lexema
	if lexema == "$" {
		lexema = "EOF"
	}
	return fmt.Sprintf(se.Message,
		lexema,
		se.Segment.Line,
	)
}
