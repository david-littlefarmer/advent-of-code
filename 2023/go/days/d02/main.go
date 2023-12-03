package d02

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
	games := parseInput(input)
	var result int
	for _, g := range games {
		if g.possibilityOfGame(12, 13, 14) {
			result += g.id
		}
	}

	return fmt.Sprint(result), time.Since(start)
}

func PartTwo(input []string) (string, time.Duration) {
	start := time.Now()
	games := parseInput(input)
	var result int
	for _, g := range games {
		result += g.minimalRequirement()
	}

	return fmt.Sprint(result), time.Since(start)
}

type game struct {
	id     int
	rounds []round
}

type round struct {
	red   int
	green int
	blue  int
}

func parseInput(input []string) []game {
	var games []game
	for _, line := range input {
		games = append(games, parseLine(line))
	}

	return games
}

func parseLine(s string) game {
	gameSplit := strings.Split(s, ":")
	id := helpers.ParseInt(strings.Fields(gameSplit[0])[1])

	g := game{
		id:     id,
		rounds: []round{},
	}

	for _, iter := range strings.Split(gameSplit[1], ";") {
		it := round{}
		for _, cubes := range strings.Split(iter, ",") {
			c := strings.Fields(cubes)
			value := helpers.ParseInt(c[0])
			switch c[1] {
			case "red":
				it.red = value
			case "green":
				it.green = value
			case "blue":
				it.blue = value
			}
		}

		g.rounds = append(g.rounds, it)
	}

	return g
}

func (gm game) possibilityOfGame(r, g, b int) bool {
	for _, round := range gm.rounds {
		if r < round.red || g < round.green || b < round.blue {
			return false
		}
	}

	return true
}

func (gm game) minimalRequirement() int {
	var r, g, b int
	for _, round := range gm.rounds {
		r = max(r, round.red)
		g = max(g, round.green)
		b = max(b, round.blue)
	}

	return r * g * b
}
