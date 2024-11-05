package varref

import (
	"fmt"
	"reflect"
	"strings"
)

type Reference struct {
	Fallback interface{}
	Path     string
}

// Todo: handle arrays as data parameter
func GetValue(data map[string]any, pathRef any) (val any) {
	ref, err := getReference(pathRef)
	if err != nil {
		return nil
	}

	value, err := traversePath(data, ref.Path)

	if err != nil {
		return ref.Fallback
	} else {
		return value
	}
}

func getReference(pathRef interface{}) (Reference, error) {
	switch path := pathRef.(type) {

	case string:
		return Reference{Path: path}, nil

	case []interface{}:
		if len(path) == 0 {
			return Reference{}, fmt.Errorf("undefined json path")
		}

		strPath, ok := path[0].(string)

		if !ok {
			return Reference{}, fmt.Errorf("invalid path value, found (%v) %v", reflect.TypeOf(path[0]), path[0])
		}

		ref := Reference{Path: strPath}

		if len(path) > 1 {
			ref.Fallback = path[1]
		}

		return ref, nil

	default:
		return Reference{}, fmt.Errorf("invalid path value, found (%v) %v", reflect.TypeOf(path), path)
	}
}

// Todo: handle array data
func traversePath(data map[string]interface{}, path string) (interface{}, error) {
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
