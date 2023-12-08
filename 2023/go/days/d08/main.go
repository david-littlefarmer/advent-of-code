package d08

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

func Run(input []string) (string, time.Duration, string, time.Duration) {
	r1, t1 := PartOne(input)
	r2, t2 := PartTwo(input)

	return r1, t1, r2, t2
}

func PartOne(input []string) (string, time.Duration) {
	start := time.Now()

	instructions, nodes := parseInput(input)
	result := part1(instructions, nodes)

	return fmt.Sprint(result), time.Since(start)
}

func PartTwo(input []string) (string, time.Duration) {
	start := time.Now()

	instructions, nodes := parseInput(input)
	result := part2(instructions, nodes)

	return fmt.Sprint(result), time.Since(start)
}

type node struct {
	l string
	r string
}

func parseInput(input []string) (instructions string, nodes map[string]node) {
	instructions = input[0]

	nodes = make(map[string]node, len(input)-2)
	for i := 2; i < len(input); i++ {
		f := func(c rune) bool {
			return !unicode.IsLetter(c)
		}

		fields := strings.FieldsFunc(input[i], f)
		n := node{
			l: fields[1],
			r: fields[2],
		}

		nodes[fields[0]] = n
	}

	return instructions, nodes
}

func part1(instructions string, nodes map[string]node) (steps int) {
	current := "AAA"
	for current != "ZZZ" {
		for _, c := range instructions {
			if c == 'L' {
				current = nodes[current].l
			} else {
				current = nodes[current].r
			}

			steps++
		}
	}

	return steps
}

func part2(instructions string, nodes map[string]node) int {
	var starters []string
	for c := range nodes {
		if c[2] == 'A' {
			starters = append(starters, c)
		}
	}

	steps := make([]int, len(starters))
	for i, current := range starters {
		for current[2] != 'Z' {
			for _, c := range instructions {
				if c == 'L' {
					current = nodes[current].l
				} else {
					current = nodes[current].r
				}

				steps[i]++
			}
		}
	}

	result := 1
	for _, s := range steps {
		result = lcm(result, s)
	}

	return result
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
