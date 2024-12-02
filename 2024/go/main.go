package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"advent-of-code/internal"
)

func main() {
	d := flag.String("d", "", "day ID to execute")
	flag.Parse()

	if d == nil {
		fmt.Println("missing required input params")
		os.Exit(1)
	}

	day, err := strconv.Atoi(*d)
	if err != nil {
		fmt.Println("couldn't parse day")
		os.Exit(1)
	}

	internal.RunDay(day)
}
