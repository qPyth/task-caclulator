package main

import (
	"fmt"
	"github.com/qPyth/task-caclulator/internal/calculator"
	"github.com/qPyth/task-caclulator/internal/types/number"
	"os"
	"regexp"
	"strconv"
)

func main() {
	calc := calculator.New(mustParse())
	res, err := calc.Calculate()
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}

func mustParse() (number.Number, number.Number, string, string) {
	if len(os.Args) != 4 {
		panic("invalid input. Usage: go run . [operand] [operator] [operand]")
	}

	operator := os.Args[2]
	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		panic("invalid input, operator must be \"+, -, *, /\"")
	}

	operand1, typ1, err := toNumber(os.Args[1])
	if err != nil {
		panic(err)
	}

	operand2, typ2, err := toNumber(os.Args[3])
	if err != nil {
		panic(err)
	}

	if operand1.Value() < 1 || operand1.Value() > 10 && operand2.Value() < 1 || operand2.Value() > 10 {
		panic("invalid input, operands must be from 1 to 10")
	}
	if typ1 != typ2 {
		panic("invalid input, operator must be same")
	}
	return operand1, operand2, operator, typ1
}

func toNumber(value string) (number.Number, string, error) {
	var err error
	var num int
	var typ string
	switch isRoman(value) {
	case true:
		num = number.RomanToArabic(value)
		typ = number.RomanStr
	case false:
		num, err = strconv.Atoi(value)
		if err != nil {
			return nil, "", fmt.Errorf("invalid input, num must be arabic or roman number")
		}
		typ = number.ArabicStr
	}
	n, err := number.New(num, typ)
	if err != nil {
		return nil, "", err
	}
	return n, typ, nil
}

func isRoman(s string) bool {

	re := regexp.MustCompile(`^(?i)(M{0,3})(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`)
	return re.MatchString(s)
}
