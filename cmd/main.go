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
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])
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
	typ := number.ArabicStr

	num2, err := strconv.Atoi(os.Args[3])
	if err != nil {
		if !isRoman(os.Args[3]) {
			panic("invalid input, operand must be arabic or roman number")
		}
		if typ == number.ArabicStr {
			panic("invalid input, operands must be only arabic or only roman")
		}
	}
	operator := os.Args[2]
	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		panic("invalid input, operator must be \"+, -, *, /\"")
	}

	operand1, err := number.New(num1, typ)
	if err != nil {
		panic(err)
	}

	operand2, err := number.New(num2, typ)
	if err != nil {
		panic(err)
	}

	if operand1.Value() < 1 || operand1.Value() > 10 && operand2.Value() < 1 || operand2.Value() > 10 {
		panic("invalid input, operands must be from 1 to 10")
	}

	return operand1, operand2, operator, typ
}

func toNumber(value string) (number.Number, error) {
	isArabic := true
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		if !isRoman(os.Args[1]) {
			return nil, fmt.Errorf("invalid input, operand must be arabic or roman number")
		}
		isArabic = false
	}
}

func isRoman(s string) bool {

	re := regexp.MustCompile(`^(?i)(M{0,3})(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`)
	return re.MatchString(s)
}
