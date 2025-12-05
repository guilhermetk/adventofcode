package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, _ := readFile(false)
	dial := 50
	password := 0
	for _, line := range content {
		if line != "" {
			steps := parseLine(line)
			dial = moveDial(dial, steps)
			fmt.Println("dial moved to: ", dial)
			if dial == 0 {
				password++
			}
		}
	}
	fmt.Println("password is: ", password)
}

func moveDial(dial, steps int) int {
	movedSteps := 0
	for movedSteps < abs(steps) {
		movedSteps++
		if steps > 0 {
			if dial == 99 {
				dial = 0
				continue
			}
			dial++
		} else {
			if dial == 0 {
				dial = 99
				continue
			}
			dial--
		}
	}

	return dial
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}

	return i
}

func parseLine(l string) int {
	direction := l[:1]
	steps := l[1:]

	convStep, err := strconv.Atoi(steps)
	if err != nil {
		log.Fatalf("error converting step to integer %s", steps)
		return 0
	}

	if direction == "L" {
		convStep *= -1
	}

	return convStep
}

func readFile(testInput bool) ([]string, error) {
	var fileType string
	if testInput {
		fileType = "test"
	} else {
		fileType = "result"
	}
	fileName := fmt.Sprintf("./input_%s.txt", fileType)
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}
