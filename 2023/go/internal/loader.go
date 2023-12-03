package internal

import (
	"bufio"
	"log"
	"os"
)

func LoadInputLines(path string) (lines []string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to load input from %s", path)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("read file, %v\n", err)
	}

	return lines
}

func LoadFirstInputLine(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to load input from %s", path)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}
