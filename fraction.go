package fraction

import (
	"strconv"
	"strings"
)

type Fraction struct {
	numerator   int
	denominator int
}

func Parse(stringFraction string) Fraction {
	fractionIndex := strings.Index(stringFraction, "/")

	if fractionIndex != -1 {
		numeratorString := stringFraction[:fractionIndex]
		denominatorString := stringFraction[fractionIndex+1:]
		numerator, _ := strconv.Atoi(numeratorString)
		denominator, _ := strconv.Atoi(denominatorString)
		return Fraction{numerator, denominator}
	}

	value, _ := strconv.Atoi(stringFraction)
	return Fraction{value, 1}
}

func Evaluate(expression string) string {
	fields := strings.Fields(expression)
	operand1 := fields[0]
	operator := fields[1]
	operand2 := fields[2]

	return operand1 + operator + operand2
}
