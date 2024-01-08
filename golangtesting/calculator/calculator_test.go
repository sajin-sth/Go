package calculator_test

import (
	"errors"
	"testing"

	"github.com/sajin-sth/golangtesting/calculator"
)

var testcases = []struct {
	name          string
	expected      float64
	expectedError error
	divisor       float64
}{
	{"division", 2.0, nil, 5.0},
	{"division by negative value", -2.0, nil, -5.0},
	{"division by zero", -2.0, errors.New("divisible by zero"), 0.0},
}

func TestDivide(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			expected := tc.expected
			got, gotError := calculator.Divide(10.0, tc.divisor)

			if gotError != nil {
				if gotError != tc.expectedError {
					t.Errorf("expected error %s, got error %s", tc.expectedError.Error(), gotError.Error())
				}
			}

			if got != expected {
				t.Errorf("expected %.1f, got %.1f", expected, got)
			}
		})
	}
}
