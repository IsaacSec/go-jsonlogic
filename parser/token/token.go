package token

import (
	"github.com/IsaacSec/go-jsonlogic/util"
	log "github.com/IsaacSec/go-jsonlogic/util/logger"
)

type Kind int
type Token interface{}
type Result interface{}

const (
	Null Kind = iota + 1
	PrimitiveVal
	ReferenceVal
	Operator
	Array
	Expression
	Object
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
		log.Error("Unexpecter bool conversion: %s", err)
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
