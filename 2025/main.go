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
	dialPos := 50
	password := 0
	zeroStepCounter := 0
	for _, line := range content {
		if line != "" {
			steps := parseLine(line)
			newDialPos, zeroStep := moveDial(dialPos, steps)
			fmt.Println("dial moved to: ", newDialPos)
			zeroStepCounter += zeroStep
			dialPos = newDialPos
			if newDialPos == 0 {
				password++
			}
		}
	}
	fmt.Println("first password is: ", password)
	fmt.Println("second password", zeroStepCounter)
}

func moveDial(dialPos, steps int) (int, int) {
	movedSteps := 0
	zeroStep := 0
	for movedSteps < abs(steps) {
		movedSteps++
		if steps > 0 {
			if dialPos == 99 {
				dialPos = 0
			} else {
				dialPos++
			}
		} else {
			if dialPos == 0 {
				dialPos = 99
			} else {
				dialPos--
			}
		}
		if dialPos == 0 {
			zeroStep++
		}
	}

	return dialPos, zeroStep
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
