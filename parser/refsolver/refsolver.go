package refsolver

import (
	"reflect"
	"strings"

	log "github.com/IsaacSec/go-jsonlogic/util/logger"
)

// Todo: handle arrays as data parameter
func GetValue(data map[string]any, ref any) (val any) {
	// Todo: add panic recover for cast error
	var current string
	var next []string
	var fb any

	switch path := ref.(type) {
	case string:
		// "foo.bar"
		current, next = popRef(path)

	case []string:
		// ["foo.bar", "fallback"]
		current, next = popRef(path[0])

		if len(path) > 1 {
			fb = path[1]
		}
	default:
		log.Warn("invalid reference: %v (%v)", path, reflect.TypeOf(ref))
		return nil
	}

	if obj, found := data[current]; found {
		if len(next) == 0 {
			return obj
		} else {
			// Todo: add error handle for cast error
			value := GetValue(obj.(map[string]any), strings.Join(next, "."))

			// Check if fallback
			if value == nil {
				return fb
			} else {
				return value
			}
		}
	} else {
		return fb
	}
}

func popRef(path string) (string, []string) {
	chain := strings.Split(path, ".")

	return chain[0], chain[1:]
}
