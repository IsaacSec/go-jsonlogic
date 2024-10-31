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
		}
	}

	return res
}

func Equals(args Args) (res token.Result) {
	args.assertHavingArgs(2)

	var first, second = args.getTwoComparableArgs()

	return first == second
}

func NotEquals(args Args) (res token.Result) {
	args.assertHavingArgs(2)

	var first, second = args.getTwoComparableArgs()

	return first != second
}

func Not(args Args) (res token.Result) {
	args.assertHavingArgs(1)

	var arg = args[0].ToBool()

	return !arg
}
