package operators

import (
	"reflect"

	"github.com/IsaacSec/go-jsonlogic/parser/token"
	"github.com/IsaacSec/go-jsonlogic/util"
	"github.com/IsaacSec/go-jsonlogic/util/logger"
)

func (args Args) getTwoComparableArgs() (a0 any, a1 any) {

	first, second := args[0].Result, args[1].Result

	// Todo: implement array and object comparison, (suggested conversion to string before)

	if reflect.TypeOf(first) == reflect.TypeOf(second) {
		a0, a1 = first, second
	} else {
		if args.ContainsNumber() {
			var err error
			a0, a1, err = args.getTwoNumericArgs()

			if err != nil {
				a0, a1 = first, second
			}
		} else {
			a0, a1 = first, second
		}
	}

	return
}

func (args Args) getTwoNumericArgs() (arg0 float64, arg1 float64, err error) {
	args.assertHavingTwoArgs()

	return util.ConvertToFloat(args[0].Result, args[1].Result)
}

func (args Args) compareNumbers(compare func(float64, float64) bool) (res token.Result) {
	first, second, err := args.getTwoNumericArgs()

	if err != nil {
		logger.Error("Argument retrieve failed: %s", err)
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
