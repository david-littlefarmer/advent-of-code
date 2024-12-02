package d02

import (
	"advent-of-code/helpers"
	"fmt"
	"strings"
	"time"
)

func Run(input []string) (string, time.Duration, string, time.Duration) {
	r1, t1 := PartOne(input)
	r2, t2 := PartTwo(input)

	return r1, t1, r2, t2
}

func PartOne(input []string) (string, time.Duration) {
	start := time.Now()

	s := parseInput(input)

	var result int
	for _, l := range s {
		if isSafeReportBasic(l) {
			result++
		}
	}

	return fmt.Sprint(result), time.Since(start)
}

func PartTwo(input []string) (string, time.Duration) {
	start := time.Now()

	s := parseInput(input)

	var result int
	for _, l := range s {
		if isSafeReportBasic(l) || isSafeReportDampened(l) {
			result++
		}
	}

	return fmt.Sprint(result), time.Since(start)
}

func isSafeReportDampened(l []int) bool {
	for i := 0; i < len(l); i++ {
		t := make([]int, 0, len(l)-1)
		for j, x := range l {
			if j != i {
				t = append(t, x)
			}
		}

		if ok := isSafeReportBasic(t); ok {
			return ok
		}

	}

	return false
}

func isSafeReportBasic(l []int) bool {
	if len(l) < 2 {
		return false
	}

	mode := determineMode(l[0], l[1])
	if mode == 0 {
		return false
	}

	safe := true
	for i := 0; i < len(l)-1; i++ {
		if !isSafeNumbers(l[i], l[i+1], mode) {
			safe = false
		}
	}

	return safe
}

func determineMode(a, b int) int {
	switch {
	case a < b:
		return 1
	case a > b:
		return -1
	default:
		return 0
	}
}

func isSafeNumbers(a, b, mode int) bool {
	if mode == 1 && a > b {
		return false
	}

	if mode == -1 && a < b {
		return false
	}

	return checkDiff(a, b)
}

func checkDiff(a, b int) bool {
	abs := helpers.Abs(a - b)
	return abs >= 1 && abs <= 3
}

func parseInput(input []string) (result [][]int) {
	s := make([][]int, 0, len(input))
	for _, l := range input {
		var sx []int
		for _, x := range strings.Split(l, " ") {
			sx = append(sx, helpers.ParseInt(x))
		}

		s = append(s, sx)
	}

	return s
}
