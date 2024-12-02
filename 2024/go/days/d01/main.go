package d01

import (
	"advent-of-code/helpers"
	"fmt"
	"slices"
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

	s1, s2 := parseInput(input)

	slices.Sort(s1)
	slices.Sort(s2)

	var result int
	for i := range s1 {
		result += helpers.Abs(s1[i] - s2[i])
	}

	return fmt.Sprint(result), time.Since(start)
}

func PartTwo(input []string) (string, time.Duration) {
	start := time.Now()

	s1, s2 := parseInput(input)

	m := map[int]int{}
	for _, x := range s2 {
		m[x] += 1
	}

	var result int
	for _, x := range s1 {
		if o, ok := m[x]; ok {
			result += x * o
		}
	}

	return fmt.Sprint(result), time.Since(start)
}

func parseInput(input []string) ([]int, []int) {
	s1 := make([]int, 0, len(input))
	s2 := make([]int, 0, len(input))

	for _, l := range input {
		nums := strings.Split(l, "   ")
		s1 = append(s1, helpers.ParseInt(nums[0]))
		s2 = append(s2, helpers.ParseInt(nums[1]))
	}

	return s1, s2
}
