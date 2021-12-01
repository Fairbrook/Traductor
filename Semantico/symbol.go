package Semantico

import "github.com/Fairbrook/analizador/Utils"

var VariableType = "variable"
var FunctionType = "function"
var ExpressionType = "expression"
var ConditionalType = "conditional"
var LiteralType = "literal"
var OperatorType = "operator"

type BaseSymbol struct {
	Identifier string
	Scope      string
	Segment    Utils.Segment
}

type Symbol interface {
	GetType() string
	GetIdentifier() string
	GetSegment() Utils.Segment
	ToArray() [3]string
	GetReturnType() string
}

type Variable struct {
	Type string
	BaseSymbol
}

func (v *Variable) GetType() string {
	return VariableType
}
func (v *Variable) GetIdentifier() string {
	return v.Identifier
}
func (v *Variable) GetSegment() Utils.Segment {
	return v.Segment
}
func (v *Variable) GetReturnType() string {
	return v.Type
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
	ParamList  []Variable
	Sentences  []Symbol
	BaseSymbol
}

func (f *Function) GetType() string {
	return FunctionType
}
func (f *Function) GetIdentifier() string {
	return f.Identifier
}
func (f *Function) GetSegment() Utils.Segment {
	return f.Segment
}
func (f *Function) GetReturnType() string {
	return f.ReturnType
}
func (f *Function) ToArray() [3]string {
	return [3]string{
		f.Identifier,
		FunctionType,
		f.ReturnType,
	}
}
func (f *Function) HasReturn() bool {
	for _, sentence := range f.Sentences {
		if sentence.GetType() == ExpressionType {
			ex := sentence.(*Expression)
			if ex.Subtype == "return" {
				return true
			}
		}
	}
	return false
}

type Expression struct {
	ReturnType string
	Subtype    string
	Parts      []Symbol
}

func (ex *Expression) GetType() string {
	return ExpressionType
}
func (ex *Expression) GetIdentifier() string {
	str := ""
	for _, item := range ex.Parts {
		str += item.GetIdentifier()
	}
	return str
}
func (ex *Expression) GetSegment() Utils.Segment {
	if len(ex.Parts) == 0 {
		return Utils.Segment{}
	}
	seg := ex.Parts[0].GetSegment()
	seg.Lexema = ex.GetIdentifier()
	return seg
}
func (ex *Expression) ToArray() [3]string {
	return [3]string{
		ex.GetIdentifier(),
		ExpressionType,
		ex.ReturnType,
	}
}
func (ex *Expression) GetReturnType() string {
	return ex.ReturnType
}

type Conditional struct {
	Sentences    []Symbol
	AltSentences []Symbol
	Subtype      string
	Expression   Expression
}

func (c *Conditional) GetType() string {
	return ConditionalType
}
func (c *Conditional) GetIdentifier() string {
	return c.Expression.GetIdentifier()
}
func (c *Conditional) GetSegment() Utils.Segment {
	return c.Expression.GetSegment()
}
func (c *Conditional) ToArray() [3]string {
	return [3]string{
		c.GetIdentifier(),
		ConditionalType,
		c.Expression.ReturnType,
	}
}
func (c *Conditional) GetReturnType() string {
	return "bool"
}
func (c *Conditional) HasAlt() bool {
	return len(c.AltSentences) > 0
}

type Literal struct {
	ReturnType string
	Value      string
	Segment    Utils.Segment
}

func (l *Literal) GetType() string {
	return LiteralType
}
func (l *Literal) GetIdentifier() string {
	return l.Value
}
func (l *Literal) GetSegment() Utils.Segment {
	return l.Segment
}
func (l *Literal) ToArray() [3]string {
	return [3]string{
		l.GetIdentifier(),
		LiteralType,
		l.ReturnType,
	}
}
func (l *Literal) GetReturnType() string {
	return l.ReturnType
}

type OperatorSymbol struct {
	Operator Operator
	Segment  Utils.Segment
}

func (o *OperatorSymbol) GetType() string {
	return OperatorType
}
func (o *OperatorSymbol) GetIdentifier() string {
	return o.Segment.Lexema
}
func (o *OperatorSymbol) GetSegment() Utils.Segment {
	return o.Segment
}
func (o *OperatorSymbol) ToArray() [3]string {
	return [3]string{
		o.GetIdentifier(),
		OperatorType,
		"",
	}
}
func (o *OperatorSymbol) GetReturnType() string {
	return ""
}
