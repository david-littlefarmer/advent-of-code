package d07_test

import (
	"testing"

	"advent-of-code/days/d07"
	"advent-of-code/internal"
)

func TestPartOne(t *testing.T) {
	in := internal.LoadInputLines("1_in_test.txt")
	out := internal.LoadFirstInputLine("1_out_test.txt")

	res, _ := d07.PartOne(in)
	if res != out {
		t.Errorf("expected result was %s, but got %s instead\n", out, res)
	}
}

func TestPartTwo(t *testing.T) {
	in := internal.LoadInputLines("2_in_test.txt")
	out := internal.LoadFirstInputLine("2_out_test.txt")

	res, _ := d07.PartTwo(in)
	if res != out {
		t.Errorf("expected result was %s, but got %s instead\n", out, res)
	}
}
