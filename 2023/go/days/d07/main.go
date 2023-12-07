package d07

import (
	"fmt"
	"sort"
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

	hands := parseInput(input)
	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j], cardsOrder1, false)
	})

	var result int
	for i, h := range hands {
		result += h.bid * (i + 1)
	}

	return fmt.Sprint(result), time.Since(start)
}

func PartTwo(input []string) (string, time.Duration) {
	start := time.Now()

	hands := parseInput(input)
	sort.Slice(hands, func(i, j int) bool {
		return compareHands(hands[i], hands[j], cardsOrder2, true)
	})

	var result int
	for i, h := range hands {
		result += h.bid * (i + 1)
	}

	return fmt.Sprint(result), time.Since(start)
}

type hand struct {
	cards string
	bid   int
}

type Strength int

const (
	FiveOfAKind Strength = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

var (
	cardsOrder1 = "AKQJT98765432"
	cardsOrder2 = "AKQT98765432J"
)

func parseInput(input []string) (hands []hand) {
	for _, s := range input {
		parts := strings.Fields(s)
		h := hand{
			cards: parts[0],
			bid:   helpers.ParseInt(parts[1]),
		}

		hands = append(hands, h)
	}

	return hands
}

func compareHands(h1, h2 hand, cardsOrder string, joker bool) bool {
	s1 := calculateStrength(h1, joker)
	s2 := calculateStrength(h2, joker)

	if s1 != s2 {
		return s1 > s2
	}

	for i := 0; i < 5; i++ {
		i1 := strings.IndexByte(cardsOrder, h1.cards[i])
		i2 := strings.IndexByte(cardsOrder, h2.cards[i])

		if i1 != i2 {
			return i1 > i2
		}

	}

	return true
}

func calculateStrength(h hand, joker bool) (st Strength) {
	cardCounts := make(map[rune]int)
	for _, card := range h.cards {
		if (joker && card != 'J') || (!joker) {
			cardCounts[card]++
		}
	}

	var maxCount int
	uniqueCards := len(cardCounts)
	for _, count := range cardCounts {
		if count > maxCount {
			maxCount = count
		}
	}

	if joker {
		for _, card := range h.cards {
			if card == 'J' {
				maxCount++
			}
		}
	}

	switch maxCount {
	case 5:
		st = FiveOfAKind
	case 4:
		st = FourOfAKind
	case 3:
		st = ThreeOfAKind
		if uniqueCards == 2 {
			st = FullHouse
		}
	case 2:
		st = OnePair
		if uniqueCards == 3 {
			st = TwoPair
		}
	case 1:
		st = HighCard
	}

	return st
}
