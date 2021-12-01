package Assembler

import "github.com/Fairbrook/analizador/Semantico"

func getSubScope(scope Semantico.Symbol, sym Semantico.Symbol) string {
	return scope.GetIdentifier() + "_" + sym.GetIdentifier()
}
