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
	mixedIndex := strings.Index(stringFraction, "_")
	fractionIndex := strings.Index(stringFraction, "/")
	whole := 0
	numerator := 0
	denominator := 1

	if (mixedIndex == -1) && (fractionIndex == -1) {
		whole, _ = strconv.Atoi(stringFraction)
	}

	if mixedIndex != -1 {
		wholeString := stringFraction[:mixedIndex]
		whole, _ = strconv.Atoi(wholeString)

		if fractionIndex != -1 {
			stringFraction = stringFraction[mixedIndex+1:]
			fractionIndex = strings.Index(stringFraction, "/")
		}
	}

	if fractionIndex != -1 {
		numeratorString := stringFraction[:fractionIndex]
		denominatorString := stringFraction[fractionIndex+1:]
		numerator, _ = strconv.Atoi(numeratorString)
		denominator, _ = strconv.Atoi(denominatorString)
	}

	return Fraction{whole*denominator + numerator, denominator}
}

func Evaluate(expression string) string {
	fields := strings.Fields(expression)
	operand1 := fields[0]
	operator := fields[1]
	operand2 := fields[2]

	return operand1 + operator + operand2
}
