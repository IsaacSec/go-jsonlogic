package parser

type Kind int
type EvalResult int
type Token interface{}

const (
	Undefined EvalResult = iota
	True
	False
)

const (
	PrimitiveVal Kind = iota + 1
	ReferenceVal
	Operator
	Array
	Expression
	Object
)

type Node struct {
	Token           Token
	Kind            Kind
	CommulativeEval EvalResult
	Childrens       *[]Node
}

func (er EvalResult) ToString() string {
	switch er {
	case Undefined:
		return "Undefined"
	case True:
		return "True"
	case False:
		return "False"
	default:
		return "Undefined"
	}
}