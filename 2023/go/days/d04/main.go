package d04

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
	cards := parseInput(input)
	result := countScorePart1(cards)

	return fmt.Sprint(result), time.Since(start)
}

func PartTwo(input []string) (string, time.Duration) {
	start := time.Now()
	cards := parseInput(input)
	result := countScorePart2(cards)

	return fmt.Sprint(result), time.Since(start)
}

type card struct {
	id       int
	winning  map[int]bool
	obtained []int
}

func parseInput(input []string) []card {
	var cards []card
	for _, line := range input {
		cards = append(cards, parseLine(line))
	}

	return cards
}

func parseLine(s string) card {
	gameSplit := strings.Split(s, ": ")
	id := helpers.ParseInt(strings.Fields(gameSplit[0])[1])

	sc := card{
		id:       id,
		winning:  map[int]bool{},
		obtained: []int{},
	}

	numbers := strings.Split(gameSplit[1], " | ")
	winning := strings.Fields(numbers[0])
	for _, n := range winning {
		sc.winning[helpers.ParseInt(n)] = true
	}

	obtained := strings.Fields(numbers[1])
	for _, n := range obtained {
		sc.obtained = append(sc.obtained, helpers.ParseInt(n))
	}

	return sc
}

func countScorePart1(cards []card) (result int) {
	for _, sc := range cards {
		count := countWins(sc)
		if count > 0 {
			result += 1 << (count - 1)
		}
	}

	return result
}

func countScorePart2(cards []card) int {
	count := map[int]int{}
	result := len(cards)

	for i, card := range cards {
		count[i] += 1
		for j := i + 1; j <= i+countWins(card); j++ {
			count[j] += count[i]
			result += count[i]
		}
	}

	return result
}

func countWins(c card) (result int) {
	for _, num := range c.obtained {
		if c.winning[num] {
			result += 1
		}
	}

	return result
}
