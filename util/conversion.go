package util

import (
	"fmt"
	"strconv"
)

func toFloat(value any) (res float64, err error) {
	switch v := any(value).(type) {
	case int:
		res = float64(v)
	case float64:
		res = v
	case string:
		res, err = strconv.ParseFloat(v, 64)
	case bool:
		if v {
			res = 1.0
		} else {
			res = 0.0
		}
	default:
		return 0, fmt.Errorf("unsupported type: %T", v)
	}

	return res, err
}

func ConvertToFloat(a any, b any) (ia float64, ib float64, err error) {

	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	ia, aErr := toFloat(a)
	ib, bErr := toFloat(b)

	if aErr != nil || bErr != nil {
		err = fmt.Errorf("arg0: [%v] or arg1: [%v] cannot be converted", aErr, bErr)
	}

	return
}
