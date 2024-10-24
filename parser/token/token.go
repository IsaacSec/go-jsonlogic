package token

import (
	"github.com/IsaacSec/go-jsonlogic/util"
	"github.com/IsaacSec/go-jsonlogic/util/logger"
)

type Kind string
type Token interface{}
type Result interface{}

const (
	Null         Kind = "null"
	PrimitiveVal Kind = "val"
	ReferenceVal Kind = "ref"
	Operator     Kind = "op"
	Array        Kind = "array"
	Expression   Kind = "exp"
	Object       Kind = "obj"
)

type Node struct {
	Token     Token
	Kind      Kind
	Childrens []*Node
}

type EvalNode struct {
	Token  Token
	Kind   Kind
	Args   []*EvalNode
	Result Result
}

func (n EvalNode) ToBool() bool {
	res, err := util.ToBool(n.Result)

	if err != nil {
		logger.Error("Unexpecter bool conversion: %s", err)
		return false
	}

	return res
}

func (n EvalNode) IsNumeric() bool {

	if _, err := util.ToFloat(n.Result); err == nil {
		return true
	} else {
		return false
	}
}
