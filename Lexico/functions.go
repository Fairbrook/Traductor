package Lexico

var simbols = map[byte]byte{
	';': 1,
	',': 2,
	'(': 3,
	')': 4,
	'{': 5,
	'}': 6,
}

var numCodes = map[string]uint8{
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
	if character == ' ' || character == '\n' || character == '\t' || character == '\r' {
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

func getSpecialType(in Segment) (out Segment) {
	out = in
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
	out.State = numCodes[out.StateName]
	return
}
