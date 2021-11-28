package fraction

import (
	"testing"
)

func TestParseFraction(t *testing.T) {
	input := "1/2"
	expected := Fraction{1, 2}

	actual := Parse(input)
	if expected != actual {
		t.Fatalf("Expected %+v, got %+v", expected, actual)
	}
}

func TestParseFractionNegative(t *testing.T) {
	input := "-1/2"
	expected := Fraction{-1, 2}

	actual := Parse(input)
	if expected != actual {
		t.Fatalf("Expected %+v, got %+v", expected, actual)
	}
}

func TestParseWholeNumber0(t *testing.T) {
	input := "0"
	expected := Fraction{0, 1}

	actual := Parse(input)
	if expected != actual {
		t.Fatalf("Expected %+v, got %+v", expected, actual)
	}
}

func TestParseWholeNumberPositive(t *testing.T) {
	input := "1"
	expected := Fraction{1, 1}

	actual := Parse(input)
	if expected != actual {
		t.Fatalf("Expected %+v, got %+v", expected, actual)
	}
}

func TestParseWholeNumberNegative(t *testing.T) {
	input := "-1"
	expected := Fraction{-1, 1}

	actual := Parse(input)
	if expected != actual {
		t.Fatalf("Expected %+v, got %+v", expected, actual)
	}
}

func TestMixedNumber(t *testing.T) {
	input := "3_3/4"
	expected := Fraction{15, 4}

	actual := Parse(input)
	if expected != actual {
		t.Fatalf("Expected %+v, got %+v", expected, actual)
	}
}

func TestMixedNumberNegative(t *testing.T) {
	input := "-3_3/4"
	expected := Fraction{-15, 4}

	actual := Parse(input)
	if expected != actual {
		t.Fatalf("Expected %+v, got %+v", expected, actual)
	}
}

func TestAsStringWholeNumber0(t *testing.T) {
	input := Fraction{0, 1}
	expected := "0"

	actual := input.AsString()
	if expected != actual {
		t.Fatalf("Expected %+v, got %+v", expected, actual)
	}
}

func TestAsStringWholeNumberNegative(t *testing.T) {
	input := Fraction{-1, 1}
	expected := "-1"

	actual := input.AsString()
	if expected != actual {
		t.Fatalf("Expected %+v, got %+v", expected, actual)
	}
}

func TestAsStringFraction(t *testing.T) {
	input := Fraction{1, 2}
	expected := "1/2"

	actual := input.AsString()
	if expected != actual {
		t.Fatalf("Expected %+v, got %+v", expected, actual)
	}
}

func TestAsStringFractionNegative(t *testing.T) {
	input := Fraction{-1, 2}
	expected := "-1/2"

	actual := input.AsString()
	if expected != actual {
		t.Fatalf("Expected %+v, got %+v", expected, actual)
	}
}

func TestAsStringMixed(t *testing.T) {
	input := Fraction{9, 8}
	expected := "1_1/8"

	actual := input.AsString()
	if expected != actual {
		t.Fatalf("Expected %+v, got %+v", expected, actual)
	}
}

func TestAsStringMixedNegative(t *testing.T) {
	input := Fraction{-9, 8}
	expected := "-1_1/8"

	actual := input.AsString()
	if expected != actual {
		t.Fatalf("Expected %+v, got %+v", expected, actual)
	}
}

func TestSimplify(t *testing.T) {
	input := Fraction{6, 8}
	expected := Fraction{3, 4}

	input.Simplify()
	actual := input
	if expected != actual {
		t.Fatalf("Expected %+v, got %+v", expected, actual)
	}
}

func TestExample1(t *testing.T) {
	input := "1/2 * 3_3/4"
	expected := "1_7/8"

	actual := Evaluate(input)
	if expected != actual {
		t.Fatalf("Expected \"%s\", got \"%s\"", expected, actual)
	}
}

func TestExample2(t *testing.T) {
	input := "2_3/8 + 9/8"
	expected := "3_1/2"

	actual := Evaluate(input)
	if expected != actual {
		t.Fatalf("Expected \"%s\", got \"%s\"", expected, actual)
	}
}
