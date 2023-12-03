package d24

import (
	"fmt"
	"time"
)

func Run(input []string) (string, time.Duration, string, time.Duration) {
	r1, t1 := PartOne(input)
	r2, t2 := PartTwo(input)

	return r1, t1, r2, t2
}

func PartOne(input []string) (string, time.Duration) {
	start := time.Now()

	result := parseInput(input)

	return fmt.Sprint(result), time.Since(start)
}

func PartTwo(input []string) (string, time.Duration) {
	start := time.Now()

	result := parseInput(input)

	return fmt.Sprint(result), time.Since(start)
}

func parseInput(input []string) (result int) {
	return result
}
