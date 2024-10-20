package operators

import (
	"fmt"

	"github.com/IsaacSec/go-jsonlogic/parser/token"
	"github.com/IsaacSec/go-jsonlogic/util"
	log "github.com/IsaacSec/go-jsonlogic/util/logger"
)

func (args Args) getTwoNumericArgs() (arg0 float64, arg1 float64, err error) {

	if len(args) < 2 {
		err = fmt.Errorf("cannot evaluate expression with less than 2 arguments, given %d", len(args))
	} else {
		arg0, arg1, err = util.ConvertToFloat(args[0].Result, args[1].Result)
	}

	return arg0, arg1, err
}

func (args Args) compareNumbers(compare func(float64, float64) bool) (res token.Result) {
	first, second, err := args.getTwoNumericArgs()

	if err != nil {
		log.Error("Argument retrieve failed: %s", err)
		res = false
		return res
	}

	return compare(first, second)
}

func (args Args) ContainsNumber() bool {
	for _, a := range args {
		if a.IsNumeric() {
			return true
		}
	}

	return false
}

func LessThan(args Args) (res token.Result) {

	return args.compareNumbers(func(a0 float64, a1 float64) bool {
		return a0 < a1
	})
}

func LessOrEqualsThan(args Args) (res token.Result) {

	return args.compareNumbers(func(a0 float64, a1 float64) bool {
		return a0 <= a1
	})
}

func GreaterThan(args Args) (res token.Result) {

	return args.compareNumbers(func(a0 float64, a1 float64) bool {
		return a0 > a1
	})
}

func GreaterOrEqualsThan(args Args) (res token.Result) {

	return args.compareNumbers(func(a0 float64, a1 float64) bool {
		return a0 >= a1
	})
}
