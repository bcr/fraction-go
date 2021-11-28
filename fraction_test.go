package fraction

import (
	"testing"
)

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
