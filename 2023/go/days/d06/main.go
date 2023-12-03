package d06

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

	races := parseInput1(input)
	result := possibleWinnings(races)

	return fmt.Sprint(result), time.Since(start)
}

func PartTwo(input []string) (string, time.Duration) {
	start := time.Now()

	race := parseInput2(input)
	result := calcRace(race)

	return fmt.Sprint(result), time.Since(start)
}

type race struct {
	t int
	d int
}

func parseInput1(input []string) (races []race) {
	times := strings.Fields(strings.Split(input[0], ":")[1])
	dists := strings.Fields(strings.Split(input[1], ":")[1])

	for i := 0; i < len(times); i++ {
		r := race{
			t: helpers.ParseInt(times[i]),
			d: helpers.ParseInt(dists[i]),
		}

		races = append(races, r)
	}

	return races
}

func parseInput2(input []string) race {
	time := strings.Join(strings.Fields(strings.Split(input[0], ":")[1]), "")
	dist := strings.Join(strings.Fields(strings.Split(input[1], ":")[1]), "")

	r := race{
		t: helpers.ParseInt(time),
		d: helpers.ParseInt(dist),
	}

	return r
}

func possibleWinnings(races []race) int {
	result := 1
	for _, r := range races {
		result *= calcRace(r)
	}

	return result
}

func calcRace(r race) (result int) {
	inRange := false
	for i := 1; i < r.t; i++ {
		if i*(r.t-i) > r.d {
			result += 1
			inRange = true
		} else if inRange {
			break
		}
	}

	return result
}
