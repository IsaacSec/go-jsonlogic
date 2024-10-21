package operators

import (
	"github.com/IsaacSec/go-jsonlogic/parser/token"
)

func And(args Args) (res token.Result) {
	res = true

	for i := range args {
		arg := args[i]

		if !arg.ToBool() {
			res = false
		}
	}

	return res
}

func Or(args Args) (res token.Result) {

	if len(args) > 0 {
		res = false
	} else { // Default value on empty "and"
		res = true
	}

	for i := range args {
		arg := args[i]

		if arg.ToBool() {
			res = true
			//break
		}
	}

	return res
}

func Equals(args Args) (res token.Result) {

	args.assertHavingTwoArgs()

	var first, second = args[0].Result, args[1].Result
	// Todo: implement array and object comparison, (suggested conversion to string before)

	if args.ContainsNumber() {
		a0, a1, err := args.getTwoNumericArgs()

		if err == nil {
			first, second = a0, a1
		}
	}

	return first == second
}

func NotEquals(args Args) (res token.Result) {

	args.assertHavingTwoArgs()

	var first, second = args[0].Result, args[1].Result

	if args.ContainsNumber() {
		a0, a1, err := args.getTwoNumericArgs()

		if err == nil {
			first, second = a0, a1
		}
	}

	return first != second
}
