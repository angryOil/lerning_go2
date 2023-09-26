package tableTest

import "errors"

func DoMath(n1, n2 int, op string) (int, error) {
	switch op {
	case "+":
		return n1 + n2, nil
	case "-":
		return n1 - n2, nil
	case "*":
		return n1 * n2, nil
	case "/":
		if n2 == 0 {
			return 0, errors.New("division by zero")
		}
		return n1 / n2, nil
	default:
		return 0, errors.New("unexpected operation:" + op)
	}
}
