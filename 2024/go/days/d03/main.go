package d03

import (
	"advent-of-code/helpers"
	"fmt"
	"regexp"
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

	in := parseInput(input)
	dos := []Do{
		{
			start: 0,
			end:   len(in) - 1,
			mode:  true,
		},
	}

	result := calculateMulsInMemory(in, dos)

	return fmt.Sprint(result), time.Since(start)
}

func PartTwo(input []string) (string, time.Duration) {
	start := time.Now()

	in := parseInput(input)
	dos := findAllDos(in)
	result := calculateMulsInMemory(in, dos)

	return fmt.Sprint(result), time.Since(start)
}

func calculateMulsInMemory(in string, dos []Do) int {
	var result int
	for _, idxs := range regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`).FindAllStringIndex(in, -1) {
		nums := strings.Split(in[idxs[0]+4:idxs[1]-1], ",")
		if checkDo(dos, idxs[0]) {
			result += helpers.ParseInt(nums[0]) * helpers.ParseInt(nums[1])
		}
	}

	return result
}

func checkDo(dos []Do, index int) bool {
	for _, d := range dos {
		if index >= d.start && index <= d.end {
			return d.mode
		}
	}

	return false
}

type Do struct {
	start int
	end   int
	mode  bool
}

func findAllDos(in string) []Do {
	dos := []Do{{start: 0, mode: true}}

	var start int
	for start < len(in) {
		doIdx := strings.Index(in[start:], "do()")
		dontIdx := strings.Index(in[start:], "don't()")
		if doIdx == -1 && dontIdx == -1 {
			break
		}

		if doIdx != -1 {
			doIdx += start + 4
		}

		if dontIdx != -1 {
			dontIdx += start + 7
		}

		mode := true
		index := doIdx
		if dontIdx != -1 && (doIdx == -1 || dontIdx < doIdx) {
			index = dontIdx
			mode = false
		}

		dos = append(dos, Do{start: index, mode: mode})
		dos[len(dos)-2].end = index - 1
		start = index
	}

	dos[len(dos)-1].end = len(in) - 1

	return dos
}

func parseInput(input []string) string {
	var in string
	for _, l := range input {
		in += l
	}

	return in
}
