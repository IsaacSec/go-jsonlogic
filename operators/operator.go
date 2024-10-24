package operators

import (
	"fmt"
	"reflect"

	"github.com/IsaacSec/go-jsonlogic/parser/token"
	"github.com/IsaacSec/go-jsonlogic/util/logger"
)

type Args []*token.EvalNode
type Operator func(Args) token.Result

type pair struct {
	Value token.Token
	_Type string
}

var Operators map[string]Operator = make(map[string]Operator)

func (args Args) assertHavingTwoArgs() {
	if len(args) < 2 {
		panic(fmt.Sprintf("cannot evaluate expression with less than 2 arguments, given %d", len(args)))
	}
}

func Run(n *token.EvalNode) (res token.Result) {

	defer func() {
		if r := recover(); r != nil {
			logger.Warn("Operation failed, Token [%v]: %s", n.Token, r)

			// Default value on panic
			res = false
		}
	}()

	switch n.Token.(type) {
	case string:
		evaluate, ok := Operators[n.Token.(string)]

		if ok {
			args := n.Args
			res = evaluate(args)

			logger.Info(
				"Evaluation: %s %v -> %v",
				n.Token,
				Args(args).GetArgValueAndType(),
				res,
			)
		} else {
			logger.Error("Undefined operator %s", n.Token)

			// Todo: change with an error handler
			res = false // Default on undefined operator
		}

	default:
		logger.Error("Token [%v] with wrong type [%v]", n.Token, reflect.TypeOf(n.Token))
	}

	return res
}

func (a Args) GetArgValueAndType() (list []pair) {
	for i := range a {
		var t string
		arg := a[i]

		if arg.Result != nil {
			t = reflect.TypeOf(arg.Result).String()
		} else {
			t = "null"
		}

		list = append(list, pair{Value: arg.Result, _Type: t})
	}

	return list
}

func (p pair) String() string {

	return fmt.Sprintf("%v (%s)", p.Value, p._Type)
}

func init() {
	Operators["and"] = And
	Operators["or"] = Or
	Operators["=="] = Equals
	Operators["!="] = NotEquals
	Operators["<"] = LessThan
	Operators["<="] = LessOrEqualsThan
	Operators[">"] = GreaterThan
	Operators[">="] = GreaterOrEqualsThan
}
