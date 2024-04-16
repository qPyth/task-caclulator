package calculator

import "github.com/qPyth/task-caclulator/internal/types/number"

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
	var resInt int
	switch c.operator {
	case plus:
		resInt = c.add()
	case minus:
		resInt = c.sub()
	case multiSign:
		resInt = c.multi()
	case divSign:
		resInt = c.div()
	}
	return number.New(resInt, c.typ)
}

func (c Calculator) add() int {
	return c.operand1.Value() + c.operand2.Value()
}

func (c Calculator) sub() int {
	return c.operand1.Value() - c.operand2.Value()
}
func (c Calculator) multi() int {
	return c.operand1.Value() * c.operand2.Value()
}
func (c Calculator) div() int {

	return c.operand1.Value() / c.operand2.Value()
}
