package token

import (
	"fmt"
	"reflect"

	"github.com/IsaacSec/go-jsonlogic/util"
	log "github.com/IsaacSec/go-jsonlogic/util/logger"
)

type Kind int
type Token interface{}
type Result interface{}

type Args []*EvalNode
type pair struct {
	Value Token
	_Type string
}

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
	Args   Args
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

func (a Args) GetArgValueAndType() (list []pair) {
	for i := range a {
		arg := a[i]
		list = append(list, pair{Value: arg.Token, _Type: reflect.TypeOf(arg.Token).String()})
	}

	return list
}

func (p pair) String() string {

	return fmt.Sprintf("%v (%s)", p.Value, p._Type)
}
