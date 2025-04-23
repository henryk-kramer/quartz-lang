package parser

type Program struct {
	Scopes []*Scope
}

type Scope struct {
	Constants []*Constant
	Functions []*Function
}

type Visibility int

const (
	PRIVATE Visibility = iota
	PUBLIC
	EXTERNAL
)

type Statement struct {
}

type Constant struct {
	Visibility *Visibility
	Identifer  *Identifer
	Expression *CompileTimeExpression
}

type Type struct {
	Visibility *Visibility
}

type Function struct {
	Visibility *Visibility
	Identifer  *Identifer
	Parameters []*FunctionParameter
	ReturnType *Type
	Body       *Scope
}

type FunctionParameter struct {
	Identifer *Identifer
	Type      *Type
}

type Identifer struct {
	name string
}

type CompileTimeExpression struct {
}
