package calculator_test

import (
	"github.com/qPyth/task-caclulator/internal/calculator"
	"github.com/qPyth/task-caclulator/internal/types/number"
	"testing"
)

func TestCalculator(t *testing.T) {
	tests := []struct {
		name     string
		operand1 number.Number
		operand2 number.Number
		operator string
		typ      string
		want     int
		wantErr  bool
	}{
		// Arabic numbers
		{"Arabic: 1 + 2", number.Arabic(1), number.Arabic(2), "+", number.ArabicStr, 3, false},
		{"Arabic: 3 - 2", number.Arabic(3), number.Arabic(2), "-", number.ArabicStr, 1, false},
		{"Arabic: 4 * 5", number.Arabic(4), number.Arabic(5), "*", number.ArabicStr, 20, false},
		{"Arabic: 6 / 3", number.Arabic(6), number.Arabic(3), "/", number.ArabicStr, 2, false},
		{"Arabic: 7 + 8", number.Arabic(7), number.Arabic(8), "+", number.ArabicStr, 15, false},
		{"Arabic: 9 - 1", number.Arabic(9), number.Arabic(1), "-", number.ArabicStr, 8, false},
		{"Arabic: 2 * 10", number.Arabic(2), number.Arabic(10), "*", number.ArabicStr, 20, false},
		{"Arabic: 10 / 2", number.Arabic(10), number.Arabic(2), "/", number.ArabicStr, 5, false},
		{"Arabic: 5 + 5", number.Arabic(5), number.Arabic(5), "+", number.ArabicStr, 10, false},
		{"Arabic: 10 - 5", number.Arabic(10), number.Arabic(5), "-", number.ArabicStr, 5, false},

		// Roman numbers
		{"Roman: I + II", number.Roman("I"), number.Roman("II"), "+", number.RomanStr, 3, false},
		{"Roman: III - II", number.Roman("III"), number.Roman("II"), "-", number.RomanStr, 1, false},
		{"Roman: IV * V", number.Roman("IV"), number.Roman("V"), "*", number.RomanStr, 20, false},
		{"Roman: VI / III", number.Roman("VI"), number.Roman("III"), "/", number.RomanStr, 2, false},
		{"Roman: VII + VIII", number.Roman("VII"), number.Roman("VIII"), "+", number.RomanStr, 15, false},
		{"Roman: IX - I", number.Roman("IX"), number.Roman("I"), "-", number.RomanStr, 8, false},
		{"Roman: II * X", number.Roman("II"), number.Roman("X"), "*", number.RomanStr, 20, false},
		{"Roman: X / II", number.Roman("X"), number.Roman("II"), "/", number.RomanStr, 5, false},
		{"Roman: V + V", number.Roman("V"), number.Roman("V"), "+", number.RomanStr, 10, false},
		{"Roman: X - V", number.Roman("X"), number.Roman("V"), "-", number.RomanStr, 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := calculator.New(tt.operand1, tt.operand2, tt.operator, tt.typ)
			got, err := c.Calculate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculator.Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Value() != tt.want {
				t.Errorf("Calculator.Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
