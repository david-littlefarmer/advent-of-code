package helpers

import (
	"log"
	"strconv"
)

func ParseInt(s string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("%v %v", err, s)
	}

	return number
}
