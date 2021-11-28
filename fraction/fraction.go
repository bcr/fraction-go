package fraction

import (
	"fmt"
	"strconv"
	"strings"
)

type Fraction struct {
	numerator   int
	denominator int
}

func (fraction *Fraction) Simplify() {
	gcd := GCD(fraction.numerator, fraction.denominator)
	fraction.numerator /= gcd
	fraction.denominator /= gcd
}

func IntAbs(value int) int {
	if value < 0 {
		return value * -1
	}

	return value
}

func (fraction *Fraction) AsString() string {
	isNegative := fraction.numerator < 0
	numerator := IntAbs(fraction.numerator)
	denominator := fraction.denominator
	whole := numerator / denominator
	numerator -= whole * denominator
	returnString := ""

	if isNegative {
		returnString += "-"
	}

	if (whole != 0) || (whole == 0 && numerator == 0) {
		returnString += fmt.Sprint(whole)
	}

	if numerator > 0 {
		if whole > 0 {
			returnString += "_"
		}
		returnString += fmt.Sprintf("%d/%d", numerator, denominator)
	}

	return returnString
}

// https://go.dev/play/p/SmzvkDjYlb
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func Parse(stringFraction string) Fraction {
	negative := stringFraction[0] == '-'
	signMultiplier := 1
	if negative {
		stringFraction = stringFraction[1:]
		signMultiplier = -1
	}

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

	return Fraction{(whole*denominator + numerator) * signMultiplier, denominator}
}

func Evaluate(expression string) string {
	fields := strings.Fields(expression)
	operand1 := Parse(fields[0])
	operator := fields[1]
	operand2 := Parse(fields[2])
	result := Fraction{0, 0}

	switch operator {
	case "*":
		result.numerator = operand1.numerator * operand2.numerator
		result.denominator = operand1.denominator * operand2.denominator

	case "+":
		result.numerator = (operand1.numerator * operand2.denominator) + (operand2.numerator * operand1.denominator)
		result.denominator = operand1.denominator * operand2.denominator

	case "-":
		result.numerator = (operand1.numerator * operand2.denominator) - (operand2.numerator * operand1.denominator)
		result.denominator = operand1.denominator * operand2.denominator

	case "/":
		result.numerator = operand1.numerator * operand2.denominator
		result.denominator = operand1.denominator * operand2.numerator

	}

	result.Simplify()
	return result.AsString()
}
