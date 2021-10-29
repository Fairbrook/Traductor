package Semantico

var VariableType = "variable"
var FunctionType = "function"

type BaseSymbol struct {
	Identifier string
	Scope      string
}

type Symbol interface {
	getType() string
	getIdentifier() string
	setScope(string)
	ToArray() [3]string
}

type Variable struct {
	Type string
	BaseSymbol
}

func (v *Variable) getType() string {
	return VariableType
}
func (v *Variable) getIdentifier() string {
	return v.Identifier
}
func (v *Variable) setScope(scope string) {
	v.Scope = scope
}
func (v *Variable) ToArray() [3]string {
	return [3]string{
		v.Identifier,
		VariableType,
		v.Type,
	}
}

type Function struct {
	ReturnType string
	Parameters map[string]Variable
	BaseSymbol
}

func (f *Function) getType() string {
	return FunctionType
}
func (f *Function) getIdentifier() string {
	return f.Identifier
}
func (f *Function) setScope(scope string) {
	f.Scope = scope
}
func (f *Function) ToArray() [3]string {
	return [3]string{
		f.Identifier,
		FunctionType,
		f.ReturnType,
	}
}
