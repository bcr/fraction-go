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

func TestExample1(t *testing.T) {
	t.Skip("Not ready yet")

	input := "1/2 * 3_3/4"
	expected := "1_7/8"

	actual := Evaluate(input)
	if expected != actual {
		t.Fatalf("Expected \"%s\", got \"%s\"", expected, actual)
	}
}

func TestExample2(t *testing.T) {
	t.Skip("Not ready yet")

	input := "2_3/8 + 9/8"
	expected := "3_1/2"

	actual := Evaluate(input)
	if expected != actual {
		t.Fatalf("Expected \"%s\", got \"%s\"", expected, actual)
	}
}
