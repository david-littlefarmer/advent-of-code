package d13

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

	floorIsLava := parseInput(input)
	result := part1(floorIsLava)

	return fmt.Sprint(result), time.Since(start)
}

func PartTwo(input []string) (string, time.Duration) {
	start := time.Now()

	floorIsLava := parseInput(input)
	result := part2(floorIsLava)

	return fmt.Sprint(result), time.Since(start)
}

func parseInput(input []string) (floorIsLava [][][]bool) {
	var valley [][]bool
	for i, row := range input {
		if row == "" {
			floorIsLava = append(floorIsLava, valley)
			valley = make([][]bool, 0)
			continue
		}

		var c []bool
		for _, col := range row {
			c = append(c, col == '#')
		}

		valley = append(valley, c)

		if i == len(input)-1 {
			floorIsLava = append(floorIsLava, valley)
		}
	}

	return floorIsLava
}

func part1(floorIsLava [][][]bool) (result int) {
	for _, valley := range floorIsLava {
		result += processValley(valley, 0)
	}

	return result
}

func part2(floorIsLava [][][]bool) (result int) {
	for _, valley := range floorIsLava {
		previous := processValley(valley, 0)
		res := 0

		for i := 0; i < len(valley) && res == 0; i++ {
			for j := 0; j < len(valley[0]) && res == 0; j++ {
				temp := deepCopyValley(valley)
				temp[i][j] = !valley[i][j]

				r := processValley(temp, previous)
				if r > 0 && r != previous {
					res = r
				}
			}
		}

		result += res
	}

	return result
}

func processValley(valley [][]bool, original int) (result int) {
	results := rows(valley)
	result = processResults(results, original, 1)

	if result == 0 || result == original {
		results := rows(rotateValley(deepCopyValley(valley)))
		result = processResults(results, original, 100)
	}

	return result
}

func processResults(results []int, original, multiplier int) (result int) {
	switch len(results) {
	case 1:
		result = results[0] * multiplier
	case 2:
		result = results[0] * multiplier
		if result == original {
			result = results[1] * multiplier
		}
	}

	return result
}

func rows(valley [][]bool) (result []int) {
	var middles []map[int]bool
	for _, row := range valley {
		middles = append(middles, mirrorRow(row))
	}

	counts := make(map[int]int, 0)
	for _, m := range middles {
		for k := range m {
			counts[k]++
		}
	}

	for k, v := range counts {
		if v == len(valley) {
			result = append(result, k)
		}
	}

	return result
}

func mirrorRow(row []bool) map[int]bool {
	middles := make(map[int]bool, 0)

	for i := 0; i < len(row)-1; i++ {
		mirrored := true
		for j := 0; i-j >= 0 && i+j+1 < len(row); j++ {
			if row[i-j] != row[i+j+1] {
				mirrored = false
				break
			}
		}

		if mirrored {
			middles[i+1] = true
		}
	}

	return middles
}

func rotateValley(valley [][]bool) [][]bool {
	rows := len(valley)
	cols := len(valley[0])

	rotated := make([][]bool, cols)
	for i := range rotated {
		rotated[i] = make([]bool, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[cols-j-1][i] = valley[i][j]
		}
	}

	return rotated
}

func deepCopyValley(valley [][]bool) [][]bool {
	copied := make([][]bool, len(valley))
	for i := range valley {
		copied[i] = make([]bool, len(valley[i]))
		copy(copied[i], valley[i])
	}

	return copied
}

func printValley(valley [][]bool) {
	for _, row := range valley {
		for _, value := range row {
			if value {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}

	fmt.Println()
}
