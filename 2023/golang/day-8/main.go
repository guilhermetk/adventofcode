package main

import (
	"fmt"
	"os"
	"strings"
)

type Coordinate struct {
	left  string
	right string
}

func main() {
	file, _ := os.ReadFile("prod.txt")
	blocks := strings.Split(string(file), "\n\n")
	directions := strings.Split(blocks[0], "")
	coordinates := strings.Split(blocks[1], "\n")
	fromTo := map[string]Coordinate{}
	for _, coordinate := range coordinates {
		if len(coordinate) > 0 {
			sides := strings.Split(coordinate, "=")
			leftSide := strings.Trim(sides[0], " ")
			rightSide := strings.Split(sides[1], ",")
			fromTo[leftSide] = Coordinate{left: strings.Trim(rightSide[0][2:], " "), right: strings.Trim(rightSide[1][:len(rightSide[1])-1], " ")}
		}
	}
	// task1(directions, fromTo)
	task2(directions, fromTo)
}

func task1(directions []string, fromTo map[string]Coordinate) {
	currentPosition := "AAA"
	stepCount := 0
	directionCount := 0
	for currentPosition != "ZZZ" {
		if directions[directionCount] == "R" {
			currentPosition = fromTo[currentPosition].right
		} else {
			currentPosition = fromTo[currentPosition].left
		}
		if directionCount+1 == len(directions) {
			directionCount = 0
		} else {
			directionCount += 1
		}
		stepCount++
		fmt.Println(stepCount, directionCount, currentPosition)
	}
}

func task2(directions []string, fromTo map[string]Coordinate) {
	currentPositions := []string{}
	for key := range fromTo {
		endWith := key[2:]
		if endWith == "A" {
			currentPositions = append(currentPositions, key)
		}
	}
	allZZZReched := false
	stepCount := 0
	directionCount := 0
	for !allZZZReched {

		endsZZZCount := 0
		for i := 0; i < len(currentPositions); i++ {
			if directions[directionCount] == "R" {
				currentPositions[i] = fromTo[currentPositions[i]].right
			} else {
				currentPositions[i] = fromTo[currentPositions[i]].left
			}
			if currentPositions[i][2:] == "Z" {
				endsZZZCount++
			}
		}

		if directionCount+1 == len(directions) {
			directionCount = 0
		} else {
			directionCount += 1
		}
		stepCount++
		if endsZZZCount == len(currentPositions) {
			allZZZReched = true
		}
	}
	fmt.Println(stepCount)
}
