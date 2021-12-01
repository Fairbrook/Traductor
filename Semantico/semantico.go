package Semantico

import (
	"github.com/Fairbrook/analizador/Sintactico"
	"github.com/Fairbrook/analizador/Utils"
)

func def(tree *Utils.Tree, table *Table) ([]error, []Symbol) {
	if tree.Children == nil {
		return make([]error, 0), make([]Symbol, 0)
	}
	defType := tree.Root.(Utils.Node).Segment.Lexema
	switch defType {
	case "Definiciones":
		{
			errors, simbols := def(tree.Children, table)
			nextErrors, nextSimbols := def(tree.Children.Next, table)
			errors = append(errors, nextErrors...)
			simbols = append(simbols, nextSimbols...)
			return errors, simbols
		}
	case "Definicion":
		return def(tree.Children, table)
	case "DefVar":
		return defVar(tree, table, "")
	case "DefFunc":
		return defFunc(tree, table)
	default:
		return make([]error, 0), make([]Symbol, 0)

	}
}

func PrintSPrototype() Function {
	return Function{
		ReturnType: "void",
		Parameters: map[string]Variable{
			"str": {Type: "char*"},
		},
		ParamList: []Variable{{Type: "char*"}},
		Sentences: []Symbol{},
		BaseSymbol: BaseSymbol{
			Identifier: "printS",
		},
	}
}

func PrintIPrototype() Function {
	return Function{
		ReturnType: "void",
		Parameters: map[string]Variable{
			"i": {Type: "int"},
		},
		ParamList: []Variable{{Type: "int"}},
		Sentences: []Symbol{},
		BaseSymbol: BaseSymbol{
			Identifier: "printI",
		},
	}
}

func PrintFPrototype() Function {
	return Function{
		ReturnType: "void",
		Parameters: map[string]Variable{
			"f": {Type: "float"},
		},
		ParamList: []Variable{{Type: "float"}},
		Sentences: []Symbol{},
		BaseSymbol: BaseSymbol{
			Identifier: "printF",
		},
	}
}

func Analize(str string) (table Table, errs []error, sentences []Symbol) {
	var err error
	var tree Utils.Tree
	table = Table{}
	printS := PrintSPrototype()
	printF := PrintFPrototype()
	printI := PrintIPrototype()
	table.Set(&printS, nil)
	table.Set(&printF, nil)
	table.Set(&printI, nil)
	if tree, err = Sintactico.GenerateSyntacticTree(str); err != nil {
		errs = []error{err}
		return
	}
	if tree.Children == nil {
		return
	}
	errs, sentences = def(tree.Children, &table)
	return
}
