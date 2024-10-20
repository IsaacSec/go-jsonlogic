package operators

import (
	"github.com/IsaacSec/go-jsonlogic/parser/token"
	log "github.com/IsaacSec/go-jsonlogic/util/logger"
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

	if len(args) < 2 {
		log.Error("Cannot evaluate expression with less than 2 arguments, given %d", len(args))
		res = false
	} else {

		var first, second = args[0].Result, args[1].Result

		// Todo: implement array and object comparison, (suggested conversion to string before)

		if args.ContainsNumber() {
			a0, a1, err := args.getTwoNumericArgs()

			if err == nil {
				first, second = a0, a1
			}
		}

		res = first == second
	}

	return res
}

func NotEquals(args Args) (res token.Result) {

	if len(args) < 2 {
		log.Error("Cannot evaluate expression with less than 2 arguments, given %d", len(args))
		res = false
	} else {

		first := args[0].Result
		second := args[1].Result

		res = first != second
	}

	return res
}
