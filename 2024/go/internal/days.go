package internal

import (
	"fmt"
	"time"

	"advent-of-code/c"
	"advent-of-code/days/d01"
	"advent-of-code/days/d02"
	"advent-of-code/days/d03"
	"advent-of-code/days/d04"
	"advent-of-code/days/d05"
	"advent-of-code/days/d06"
	"advent-of-code/days/d07"
	"advent-of-code/days/d08"
	"advent-of-code/days/d09"
	"advent-of-code/days/d10"
	"advent-of-code/days/d11"
	"advent-of-code/days/d12"
	"advent-of-code/days/d13"
	"advent-of-code/days/d14"
	"advent-of-code/days/d15"
	"advent-of-code/days/d16"
	"advent-of-code/days/d17"
	"advent-of-code/days/d18"
	"advent-of-code/days/d19"
	"advent-of-code/days/d20"
	"advent-of-code/days/d21"
	"advent-of-code/days/d22"
	"advent-of-code/days/d23"
	"advent-of-code/days/d24"
	"advent-of-code/days/d25"
)

func RunDay(day int) {
	filepath := fmt.Sprintf("inputs/input_%02d.txt", day)
	input := LoadInputLines(filepath)
	days := map[int]func([]string) (string, time.Duration, string, time.Duration){
		1:  d01.Run,
		2:  d02.Run,
		3:  d03.Run,
		4:  d04.Run,
		5:  d05.Run,
		6:  d06.Run,
		7:  d07.Run,
		8:  d08.Run,
		9:  d09.Run,
		10: d10.Run,
		11: d11.Run,
		12: d12.Run,
		13: d13.Run,
		14: d14.Run,
		15: d15.Run,
		16: d16.Run,
		17: d17.Run,
		18: d18.Run,
		19: d19.Run,
		20: d20.Run,
		21: d21.Run,
		22: d22.Run,
		23: d23.Run,
		24: d24.Run,
		25: d25.Run,
	}

	if day > 0 {
		r1, t1, r2, t2 := days[day](input)
		fmt.Printf("%s %s\n\n%s - %s\n%s\n\n%s - %s\n%s\n\n",
			c.Background(" File ", c.FgBlack, c.BgMagenta),
			filepath,
			c.Background(" Part 1 ", c.FgBlack, c.BgGreen),
			c.Foreground(t1.String(), c.FgBlue),
			c.Foreground(r1, c.FgGreen),
			c.Background(" Part 2 ", c.FgBlack, c.BgYellow),
			c.Foreground(t2.String(), c.FgBlue),
			c.Foreground(r2, c.FgYellow),
		)
	}
}
