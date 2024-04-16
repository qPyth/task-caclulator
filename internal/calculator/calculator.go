package calculator

import (
	"fmt"
	"github.com/qPyth/task-caclulator/internal/types/number"
)

var (
	plus      = "+"
	minus     = "-"
	multiSign = "*"
	divSign   = "/"
)

type Calculator struct {
	operand1 number.Number
	operand2 number.Number
	operator string
	typ      string
}

func New(operand1 number.Number, operand2 number.Number, operator string, typ string) *Calculator {
	return &Calculator{operand1: operand1, operand2: operand2, operator: operator, typ: typ}
}

func (c Calculator) Calculate() (number.Number, error) {
	var err error
	var resInt int
	switch c.operator {
	case plus:
		resInt = c.add()
	case minus:
		resInt, err = c.sub()
		if err != nil {
			return nil, err
		}
	case multiSign:
		resInt = c.multi()
	case divSign:
		resInt, err = c.div()
		if err != nil {
			return nil, err
		}
	}
	return number.New(resInt, c.typ)
}

func (c Calculator) add() int {
	return c.operand1.Value() + c.operand2.Value()
}

func (c Calculator) sub() (int, error) {
	res := c.operand1.Value() - c.operand2.Value()
	if c.typ == number.RomanStr && res <= 0 {
		return 0, fmt.Errorf("result is negative or equal 0")
	}
	return res, nil
}
func (c Calculator) multi() int {
	return c.operand1.Value() * c.operand2.Value()
}
func (c Calculator) div() (int, error) {
	if c.operand2.Value() == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return c.operand1.Value() / c.operand2.Value(), nil
}
