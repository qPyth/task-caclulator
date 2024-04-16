package number

import "fmt"

const (
	RomanStr  = "roman"
	ArabicStr = "arabic"
)

type Arabic int

func (a Arabic) Value() int {
	return int(a)
}

type Number interface {
	Value() int
}

type Roman string

func (r Roman) Value() int {
	return RomanToArabic(string(r))
}

func New(value int, numType string) (Number, error) {
	switch numType {
	case ArabicStr:
		return Arabic(value), nil
	case RomanStr:
		return Roman(ArabicToRoman(value)), nil
	default:
		return nil, fmt.Errorf("unknown number type: %s", numType)
	}
}
