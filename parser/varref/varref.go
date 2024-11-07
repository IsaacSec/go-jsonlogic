package varref

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/IsaacSec/go-jsonlogic/util"
)

type Reference struct {
	Fallback interface{}
	Path     string
	Index    int
}

func GetValue(data any, pathRef any) (val any) {
	ref, err := getReference(pathRef)
	if err != nil {
		return nil
	}

	value, err := traverseReference(data, ref)

	if err != nil {
		return ref.Fallback
	} else {
		return value
	}
}

func getReference(pathRef interface{}) (Reference, error) {
	switch path := pathRef.(type) {
	case float64:
		index, err := util.ToInt(path)
		return Reference{Index: index}, err
	case int:
		return Reference{Index: path}, nil

	case string:
		return Reference{Path: path}, nil

	case []interface{}:
		ref := Reference{}

		if len(path) == 0 {
			return ref, fmt.Errorf("undefined json path")
		}

		if len(path) > 1 {
			ref.Fallback = path[1]
		}

		if index, err := util.ToInt(path[0]); err == nil {
			ref.Index = index
		} else {
			if strPath, ok := path[0].(string); ok {
				ref.Path = strPath
			} else {
				return ref, fmt.Errorf("invalid path value, found (%v) %v", reflect.TypeOf(path[0]), path[0])
			}
		}

		return ref, nil

	default:
		return Reference{}, fmt.Errorf("invalid path value, found (%v) %v", reflect.TypeOf(path), path)
	}
}

// Todo: handle array data
func traverseReference(data any, ref Reference) (interface{}, error) {
	switch d := data.(type) {
	case []interface{}:
		return getFromArray(d, ref.Index), nil

	case interface{}:
		if input, ok := d.(map[string]any); !ok {
			return nil, fmt.Errorf("invalid input of type=%v", reflect.TypeOf(d))
		} else {
			return getFromMap(input, ref.Path)
		}

	default:
		return nil, fmt.Errorf("invalid input of type=%v", reflect.TypeOf(d))
	}
}

func getFromMap(data map[string]any, path string) (interface{}, error) {
	keys := strings.Split(path, ".")
	current := data

	for i, key := range keys {
		value, exists := current[key]

		if !exists {
			return nil, fmt.Errorf("key not found: %s", key)
		}

		// Last path reference found
		if i >= (len(keys) - 1) {
			return value, nil
		}

		// ensure the value is a map and continue
		next, ok := value.(map[string]interface{})

		if !ok {
			return nil, fmt.Errorf("expected map but found: %v", value)
		}

		current = next
	}

	// Path is empty, the whole data must be returned
	return data, nil
}

func getFromArray(data []any, index int) interface{} {
	if index >= len(data) || index < 0 { // Index out of bound
		return nil
	}

	return data[index]
}
