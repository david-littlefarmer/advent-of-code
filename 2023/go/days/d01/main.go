package d01

import (
	"fmt"
	"strings"
	"time"
	"unicode"

	"advent-of-code/helpers"
)

func Run(input []string) (string, time.Duration, string, time.Duration) {
	r1, t1 := PartOne(input)
	r2, t2 := PartTwo(input)

	return r1, t1, r2, t2
}

func PartOne(input []string) (string, time.Duration) {
	start := time.Now()
	var result int
	for _, line := range input {
		result += parseInput(line, false)
	}

	return fmt.Sprint(result), time.Since(start)
}

func PartTwo(input []string) (string, time.Duration) {
	start := time.Now()
	var result int
	for _, line := range input {
		result += parseInput(line, true)
	}

	return fmt.Sprint(result), time.Since(start)
}

func parseInput(s string, parseStrings bool) int {
	digits := make([]int, len(s))
	for i, r := range s {
		if unicode.IsDigit(r) {
			digits[i] = helpers.ParseInt(string(r))
		}
	}

	if parseStrings {
		parseStringNumbers(s, digits)
	}

	return firstNonZero(digits)*10 + lastNonZero(digits)
}

func parseStringNumbers(s string, digits []int) {
	for i, sd := range stringDigits {
		if index := strings.Index(s, sd); index != -1 {
			digits[index] = i + 1
		}

		if index := strings.LastIndex(s, sd); index != -1 {
			digits[index] = i + 1
		}
	}
}

func firstNonZero(slice []int) int {
	for _, value := range slice {
		if value != 0 {
			return value
		}
	}

	return -1
}

func lastNonZero(slice []int) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] != 0 {
			return slice[i]
		}
	}

	return -1
}

var stringDigits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
