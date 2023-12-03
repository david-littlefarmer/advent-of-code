package d03

import (
	"fmt"
	"strings"
	"time"
	"unicode"

	"advent-of-code/helpers"
)

func Run(input []string) (string, time.Duration, string, time.Duration) {
	r1, t1 := PartOne(input)
	r2, t2 := PartTwo(input)

	return r1, t1, r2, t2
}

func PartOne(input []string) (string, time.Duration) {
	start := time.Now()
	schematic := parseInput(input)
	result := schematicSum(schematic)

	return fmt.Sprint(result), time.Since(start)
}

func PartTwo(input []string) (string, time.Duration) {
	start := time.Now()
	schematic := parseInput(input)
	result := schematicGearRatios(schematic)

	return fmt.Sprint(result), time.Since(start)
}

func parseInput(input []string) [][]rune {
	schematic := make([][]rune, 0, len(input)+2)
	blankLine := []rune(strings.Repeat(".", len(input[0])+2))

	schematic = append(schematic, blankLine)
	for _, line := range input {
		schematic = append(schematic, append([]rune{'.'}, append([]rune(line), '.')...))
	}
	schematic = append(schematic, blankLine)

	return schematic
}

func schematicSum(schematic [][]rune) (result int) {
	var numString string
	var inNum, isPartNumber bool
	for i, l := range schematic {
		for j, r := range l {
			if unicode.IsDigit(r) {
				inNum = true
				numString += string(r)
				for x := i - 1; x <= i+1; x++ {
					for y := j - 1; y <= j+1; y++ {
						if isSymbol(schematic[x][y]) {
							isPartNumber = true
						}
					}
				}
			} else if inNum {
				if isPartNumber {
					result += helpers.ParseInt(numString)
					isPartNumber = false
				}

				inNum = false
				numString = ""
			}
		}
	}

	return result
}

func isSymbol(r rune) bool {
	return !unicode.IsDigit(r) && r != '.'
}

func schematicGearRatios(schematic [][]rune) (result int) {
	var numString string
	var first, second int
	for i, line := range schematic {
		for j, r := range line {
			if r == '*' {
				for x := i - 1; x <= i+1; x++ {
					for y := j - 1; y <= j+1; y++ {
						if unicode.IsDigit(schematic[x][y]) {
							yy := y
							for ; unicode.IsDigit(schematic[x][yy]); yy-- {
							}

							for yy++; unicode.IsDigit(schematic[x][yy]); yy++ {
								numString += string(schematic[x][yy])
								y = yy + 1
							}

							first = second
							second = helpers.ParseInt(numString)
						}

						numString = ""
					}
				}
				result += first * second
				first, second = 0, 0
			}
		}
	}

	return result
}
