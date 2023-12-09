package d09

import (
	"fmt"
	"strings"
	"time"

	"advent-of-code/helpers"
)

func Run(input []string) (string, time.Duration, string, time.Duration) {
	r1, t1 := PartOne(input)
	r2, t2 := PartTwo(input)

	return r1, t1, r2, t2
}

func PartOne(input []string) (string, time.Duration) {
	start := time.Now()

	reports := parseInput(input)
	var result int
	for _, r := range reports {
		result += processReport(reverseSlice(r))
	}

	return fmt.Sprint(result), time.Since(start)
}

func PartTwo(input []string) (string, time.Duration) {
	start := time.Now()

	reports := parseInput(input)
	var result int
	for _, r := range reports {
		result += processReport(r)
	}
	return fmt.Sprint(result), time.Since(start)
}

func parseInput(input []string) (reports [][]int) {
	for _, s := range input {
		nums := strings.Fields(s)
		report := make([]int, 0, len(nums))
		for _, n := range nums {
			report = append(report, helpers.ParseInt(n))
		}

		reports = append(reports, report)
	}

	return reports
}

func reverseSlice(s []int) []int {
	length := len(s)
	reversed := make([]int, length)

	for i := 0; i < length; i++ {
		reversed[i] = s[length-i-1]
	}

	return reversed
}

func processReport(report []int) (result int) {
	lines := buildTree(report)

	first := lines[len(lines)-1][0]
	for i := len(lines) - 2; i >= 0; i-- {
		first = lines[i][0] - first
	}

	return first
}

func buildTree(report []int) [][]int {
	lines := [][]int{report}
	current := lines[0]
	for !allZeroes(current) {
		next := make([]int, 0, len(current)-1)
		for i := 0; i < len(current)-1; i++ {
			next = append(next, (current[i+1] - current[i]))
		}

		current = next
		lines = append(lines, current)
	}

	return lines
}

func allZeroes(s []int) bool {
	for _, v := range s {
		if v != 0 {
			return false
		}
	}

	return true
}
