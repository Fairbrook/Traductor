package Semantico

type Operator struct {
	isBinary      bool
	acceptedTypes []string
	returnType    string
	name          string
}

func (op *Operator) acceptType(t string) bool {
	for _, k := range op.acceptedTypes {
		if k == t {
			return true
		}
	}
	return false
}

var OpByType = map[int]Operator{
	5: {
		isBinary:      true,
		acceptedTypes: []string{"int", "float", "char*"},
		returnType:    "same",
		name:          "suma",
	},
	6: {
		isBinary:      true,
		acceptedTypes: []string{"int", "float"},
		returnType:    "same",
		name:          "mul",
	},
	7: {
		isBinary:      true,
		acceptedTypes: []string{"int", "float", "char*"},
		returnType:    "bool",
		name:          "relac",
	},
	8: {
		isBinary:      true,
		acceptedTypes: []string{"bool"},
		returnType:    "bool",
		name:          "or",
	},
	9: {
		isBinary:      true,
		acceptedTypes: []string{"bool"},
		returnType:    "bool",
		name:          "and",
	},
	10: {
		isBinary:      false,
		acceptedTypes: []string{"bool"},
		returnType:    "same",
		name:          "not",
	},
	11: {
		isBinary:      true,
		acceptedTypes: []string{"bool", "float", "int", "char*"},
		returnType:    "bool",
		name:          "igual",
	},
}
