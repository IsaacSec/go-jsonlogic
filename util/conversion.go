package util

import (
	"fmt"
	"strconv"
)

type Numeric interface {
	int | float64 | string | bool
}

func toFloat[T Numeric](value T) (res float64, err error) {
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

func ToFloat[T Numeric, V Numeric](a T, b V) (ia float64, ib float64, err error) {

	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	ia, aErr := toFloat(a)
	ib, bErr := toFloat(b)

	if aErr != nil || bErr != nil {
		err = fmt.Errorf("conversion error with arg0: [%v] arg1: [%v]", aErr, bErr)
	}

	return
}
