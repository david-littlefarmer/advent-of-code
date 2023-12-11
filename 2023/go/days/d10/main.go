package d10

import (
	"fmt"
	"strconv"
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

	pm, x, y := parseInput(input)
	result := part1(pm, x, y)

	return fmt.Sprint(result), time.Since(start)
}

// Bad solution, hardcoded start letter without detection
func PartTwo(input []string) (string, time.Duration) {
	start := time.Now()

	pm, x, y := parseInput(input)
	result := part2(pm, x, y)

	return fmt.Sprint(result), time.Since(start)
}

func parseInput(input []string) (pm [][]rune, x, y int) {
	for r, row := range input {
		line := make([]rune, 0, len(row))
		for c, col := range row {
			line = append(line, col)
			if col == 'S' {
				x, y = r, c
			}
		}

		pm = append(pm, line)
	}

	return pm, x, y
}

func part1(pm [][]rune, x, y int) (result int) {
	xs, ys, _ := findStartingPipes(pm, x, y)

	xp, yp := x, y
	xn, yn := xs, ys
	for pm[xn][yn] != 'S' {
		x1, y1, x2, y2, _ := nextMove(xn, yn, pm[xn][yn])
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

func findStartingPipes(pm [][]rune, xs, ys int) (xr, yr int, s rune) {
	offsets := map[string][]int{
		"N": {-1, 0},
		"S": {1, 0},
		"E": {0, 1},
		"W": {0, -1},
	}
	var dirs string
	for k, offset := range offsets {
		x, y := xs+offset[0], ys+offset[1]
		if isValidPosition(pm, x, y) {
			x1, y1, x2, y2, found := nextMove(x, y, pm[x][y])
			if found && ((xs == x1 && ys == y1) || (xs == x2 && ys == y2)) {
				dirs += k
				xr, yr = x, y
			}
		}
	}
	switch {
	case strings.Contains(dirs, "N") && strings.Contains(dirs, "S"):
		s = '|'
	case strings.Contains(dirs, "E") && strings.Contains(dirs, "W"):
		s = '-'
	case strings.Contains(dirs, "N") && strings.Contains(dirs, "E"):
		s = 'L'
	case strings.Contains(dirs, "N") && strings.Contains(dirs, "W"):
		s = 'J'
	case strings.Contains(dirs, "S") && strings.Contains(dirs, "E"):
		s = 'F'
	case strings.Contains(dirs, "S") && strings.Contains(dirs, "W"):
		s = '7'
	}

	return xr, yr, s
}

func isValidPosition(pm [][]rune, x, y int) bool {
	return x >= 0 && x < len(pm[0]) && y >= 0 && y < len(pm)
}

func nextMove(x, y int, r rune) (x1, y1, x2, y2 int, found bool) {
	switch r {
	case '-':
		x1, y1 = x, y-1
		x2, y2 = x, y+1
	case '|':
		x1, y1 = x-1, y
		x2, y2 = x+1, y
	case 'L':
		x1, y1 = x-1, y
		x2, y2 = x, y+1
	case 'J':
		x1, y1 = x-1, y
		x2, y2 = x, y-1
	case '7':
		x1, y1 = x, y-1
		x2, y2 = x+1, y
	case 'F':
		x1, y1 = x, y+1
		x2, y2 = x+1, y
	default:
		return x1, y1, x2, y2, false
	}

	// fmt.Printf("%s [%d:%d] - [%d:%d] [%d:%d]\n", string(r), x, y, x1, y1, x2, y2)
	return x1, y1, x2, y2, true
}

func part2(pm [][]rune, x, y int) (result int) {
	pipe := make([][]rune, len(pm))
	for i := range pipe {
		pipe[i] = make([]rune, len(pm[0]))
	}

	for i := range pipe {
		for j := range pipe[i] {
			pipe[i][j] = '.'
		}
	}

	xs, ys, s := findStartingPipes(pm, x, y)

	pipe[x][y] = s

	xp, yp := x, y
	xn, yn := xs, ys
	for pm[xn][yn] != 'S' {
		x1, y1, x2, y2, _ := nextMove(xn, yn, pm[xn][yn])
		xt, yt := xn, yn

		pipe[xn][yn] = pm[xn][yn]

		xn, yn = x1, y1
		if xn == xp && yn == yp {
			xn, yn = x2, y2
		}

		xp, yp = xt, yt

	}

	for i, x := range pipe {
		var in bool
		var up bool
		for j := 0; j < len(x); j++ {
			r := pipe[i][j]

			switch r {
			case '.':
				if in {
					result++
					pipe[i][j] = 'X'
				}
			case '|':
				in = !in
			case 'L':
				in = !in
				up = false
			case 'F':
				in = !in
				up = true
			case 'J':
				if !up {
					in = !in
				}
			case '7':
				if up {
					in = !in
				}
			}
		}
	}

	printit := true
	if printit {
		l := len(pipe[0])
		digits := len(strconv.Itoa(l))
		for i := digits; i > 0; i-- {
			fmt.Printf("%s", strings.Repeat(" ", digits+1))
			for j := 0; j < l; j++ {
				s := strconv.Itoa(j)
				d := len(s)
				if d >= i {
					fmt.Printf("%s", s[d-i:d-i+1])
				} else {
					fmt.Printf(" ")
				}
			}
			fmt.Printf("\n")
		}

		for i, x := range pipe {
			fmt.Printf("%3d ", i)

			for _, y := range x {
				if y == '.' {
					fmt.Printf("%s", string(y))
				} else if y == 'X' {
					fmt.Printf("\033[31m%s\033[0m", string(y))
				} else {
					fmt.Printf("\033[32m%s\033[0m", string(y))
				}
			}

			fmt.Printf("\n")
		}
	}
	return result
}
