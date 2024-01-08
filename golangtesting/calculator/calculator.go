package calculator

import "errors"

func Divide(a, b float64) (float64, error) {

	if b == 0.0 {
		return 0.0, errors.New("division by zero")
	}
	return a / b, nil
}
