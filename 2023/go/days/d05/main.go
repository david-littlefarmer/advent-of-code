package d05

import (
	"fmt"
	"strings"
	"sync"
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
	a := parseInput(input)

	result := seedToLocationPart1(a)
	return fmt.Sprint(result), time.Since(start)
}

func PartTwo(input []string) (string, time.Duration) {
	start := time.Now()
	a := parseInput(input)

	result := seedToLocationPart2(a)
	return fmt.Sprint(result), time.Since(start)
}

type almanac struct {
	seeds []int
	maps  [][]sm
}

type sm struct {
	d int
	s int
	r int
}

func parseInput(input []string) *almanac {
	a := &almanac{}
	inputSeeds := strings.Fields(strings.Split(input[0], ": ")[1])
	for _, s := range inputSeeds {
		a.seeds = append(a.seeds, helpers.ParseInt(s))
	}

	var currentMap []sm
	for i := 2; i < len(input); i++ {
		s := input[i]
		if s == "" || i == len(input)-1 {
			if currentMap != nil {
				a.maps = append(a.maps, currentMap)
			}
			currentMap = nil
			continue
		}

		if strings.HasSuffix(s, "map:") {
			currentMap = make([]sm, 0)
			continue
		}

		if currentMap != nil {
			fields := strings.Fields(s)
			seedMap := sm{
				d: helpers.ParseInt(fields[0]),
				s: helpers.ParseInt(fields[1]),
				r: helpers.ParseInt(fields[2]),
			}
			currentMap = append(currentMap, seedMap)
		}
	}

	return a
}

func parseInputSeedsPart2(input string) (seeds []int) {
	inputSeeds := strings.Fields(input)
	var is []int
	for _, s := range inputSeeds {
		is = append(is, helpers.ParseInt(s))
	}

	for i := 0; i < len(is); i += 2 {
		for j := 0; j < is[i+1]; j++ {
			seeds = append(seeds, is[i]+j)
		}
	}

	return seeds
}

func seedToLocationPart1(a *almanac) (result int) {
	result = seedWay(a, a.seeds[0])
	for _, s := range a.seeds {
		result = min(result, seedWay(a, s))
	}

	return result
}

func seedToLocationPart2(a *almanac) (result int) {
	result = seedWay(a, a.seeds[0])
	var wg sync.WaitGroup
	resultChan := make(chan int, len(a.seeds))

	for i := 0; i < len(a.seeds); i += 2 {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			res := seedWay(a, a.seeds[index])
			for j := 1; j < a.seeds[index+1]; j++ {
				res = min(res, seedWay(a, a.seeds[index]+j))
			}

			resultChan <- res
		}(i)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for val := range resultChan {
		result = min(result, val)
	}

	return result
}

func seedWay(a *almanac, s int) int {
	for _, m := range a.maps {
		for _, sm := range m {
			if s >= sm.s && s <= sm.s+sm.r {
				t := s - sm.s
				s = sm.d + t
				break
			}
		}
	}

	return s
}
