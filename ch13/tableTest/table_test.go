package tableTest

import "testing"

func TestTable(t *testing.T) {
	data := []struct {
		name     string
		num1     int
		num2     int
		op       string
		expected int
		errMsg   string
	}{
		{"add", 2, 2, "+", 4, ""},
		{"sub", 2, 2, "-", 0, ""},
		{"mul", 2, 2, "*", 4, ""},
		{"div", 2, 2, "/", 1, ""},
		{"err", 2, 0, "/", 0, "division by zero"},
		{"unexpected Op", 2, 0, "z", 0, "unexpected operation:z"},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			result, err := DoMath(d.num1, d.num2, d.op)
			if result != d.expected {
				t.Errorf("expected %d , got %d ", d.expected, result)
			}
			var errMsg string
			if err != nil {
				errMsg = err.Error()
			}
			if errMsg != d.errMsg {
				t.Errorf("expected error msg %s but got %s", d.errMsg, errMsg)
			}
		})
	}
}
