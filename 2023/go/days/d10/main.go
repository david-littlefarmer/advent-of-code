package d10

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

	pm, x, y := parseInput(input)
	result := part1(pm, x, y)

	return fmt.Sprint(result), time.Since(start)
}

func PartTwo(input []string) (string, time.Duration) {
	start := time.Now()

	_, _, _ = parseInput(input)

	return fmt.Sprint(0), time.Since(start)
}

func parseInput(input []string) (pm [][]rune, x, y int) {
	for i, l := range input {
		line := make([]rune, 0, len(l))
		for j, r := range l {
			line = append(line, r)
			if r == 'S' {
				x, y = j, i
			}
		}

		pm = append(pm, line)
	}

	return pm, x, y
}

func part1(pm [][]rune, x, y int) (result int) {
	xs, ys := findStartingPipes(pm, x, y)
	fmt.Println("starting pipes", xs, ys)

	xp, yp := xs, ys
	xn, yn := xs, ys
	for pm[yn][xn] != 'S' {
		x1, y1, x2, y2, _ := nextMove(xn, yn, pm[yn][xn])
		xt, yt := xn, yn

		xn, yn = x1, y1
		if xn == xp && yn == yp {
			xn, yn = x2, y2
		}

		xp, yp = xt, yt

		result++
	}

	return (result + 1) / 2
}

func findStartingPipes(pm [][]rune, xs, ys int) (int, int) {
	offsets := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, offset := range offsets {
		x, y := xs+offset[0], ys+offset[1]
		if isValidPosition(pm, x, y) {
			x1, y1, x2, y2, found := nextMove(x, y, pm[y][x])
			if found && ((xs == x1 && ys == y1) || (xs == x2 && ys == y2)) {
				return x, y
			}
		}
	}

	return 0, 0
}

func isValidPosition(pm [][]rune, x, y int) bool {
	return x >= 0 && x < len(pm[0]) && y >= 0 && y < len(pm)
}

func nextMove(x, y int, r rune) (x1, y1, x2, y2 int, found bool) {
	switch r {
	case '-':
		x1, y1 = x-1, y
		x2, y2 = x+1, y
	case '|':
		x1, y1 = x, y-1
		x2, y2 = x, y+1
	case 'L':
		x1, y1 = x, y-1
		x2, y2 = x+1, y
	case 'J':
		x1, y1 = x, y-1
		x2, y2 = x-1, y
	case '7':
		x1, y1 = x-1, y
		x2, y2 = x, y+1
	case 'F':
		x1, y1 = x+1, y
		x2, y2 = x, y+1
	default:
		return x1, y1, x2, y2, false
	}

	fmt.Printf("%s [%d:%d] - [%d:%d] [%d:%d]\n", string(r), x, y, x1, y1, x2, y2)
	return x1, y1, x2, y2, true
}
