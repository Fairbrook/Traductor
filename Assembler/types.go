package Assembler

import "io"

var translateTypes = map[string]string{
	"int":   "DWORD",
	"float": "REAL4",
	"char*": "byte",
}

var specialFunctions = map[string]string{
	"printS": "printf(",
	"printF": "printf(\"%f\",",
	"printI": "printf(\"%u\",",
}

type Translator struct {
	Filename      string
	tags_counter  map[string]int64
	stream        io.Writer
	fileNameNoExt string
}
