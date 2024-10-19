package operators

import (
	"github.com/IsaacSec/go-jsonlogic/parser/token"
	"github.com/IsaacSec/go-jsonlogic/util"
	log "github.com/IsaacSec/go-jsonlogic/util/logger"
)

func LessThan(args token.Args) (res token.Result) {

	if len(args) < 2 {
		log.Error("Cannot evaluate expression with less than 2 arguments, given %d", len(args))
		res = false
	} else {
		first, second, err := util.ConvertToFloat(args[0].Result, args[1].Result)

		if err != nil {
			log.Warn("Conversion failed: %s", err)
			res = false
			return res
		}

		res = first < second
	}

	return res

}
