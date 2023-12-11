package d11

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

	rows, cols := parseInput(input)
	coords := getCords(input, 2, rows, cols)

	result := countPaths(coords)

	return fmt.Sprint(result), time.Since(start)
}

func PartTwo(input []string) (string, time.Duration) {
	start := time.Now()

	rows, cols := parseInput(input)
	coords := getCords(input, 1000000, rows, cols)

	result := countPaths(coords)

	return fmt.Sprint(result), time.Since(start)
}

func parseInput(input []string) (map[int]bool, map[int]bool) {
	rows := map[int]bool{}
	cols := map[int]bool{}

	for r, row := range input {
		if !strings.Contains(row, "#") {
			rows[r] = true
		}
	}

	for r := range input[0] {
		var withGalaxy bool
		for c := range input {
			if input[c][r] == '#' {
				withGalaxy = true
			}
		}

		if !withGalaxy {
			cols[r] = true
		}

	}

	return rows, cols
}

type coord struct {
	x, y int
}

func getCords(input []string, size int, rows, cols map[int]bool) (coords []coord) {
	for r, row := range input {
		for c := range row {
			if input[r][c] == '#' {
				emptyCols := 0
				emptyRows := 0

				for i := 0; i < r; i++ {
					if rows[i] {
						emptyRows++
					}
				}

				for i := 0; i < c; i++ {
					if cols[i] {
						emptyCols++
					}
				}

				co := coord{
					x: r + emptyRows*(size-1),
					y: c + emptyCols*(size-1),
				}
				coords = append(coords, co)
			}
		}
	}

	return coords
}

func countPaths(coords []coord) (result int) {
	for i := range coords {
		for j := i + 1; j < len(coords); j++ {
			result += helpers.Abs(coords[i].x-coords[j].x) + helpers.Abs(coords[i].y-coords[j].y)
		}
	}

	return result
}
