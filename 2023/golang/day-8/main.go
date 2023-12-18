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
	steps := []int{}
	for key := range fromTo {
		endWith := key[2:]
		if endWith == "A" {
			steps = append(steps, getStepsCount(directions, key, fromTo, 0, 0))
		}
	}
	fmt.Println(LCM(steps[0], steps[1], steps[2], steps[3], steps[4], steps[5]))
}

func getStepsCount(directions []string, position string, fromTo map[string]Coordinate, stepsCount int, directionCount int) int {
	if position[2:] == "Z" {
		return stepsCount
	}

	var newPosition string
	if directions[directionCount] == "R" {
		newPosition = fromTo[position].right
	} else {
		newPosition = fromTo[position].left
	}

	var newDirection = directionCount
	if directionCount+1 == len(directions) {
		newDirection = 0
	} else {
		newDirection += 1
	}
	return getStepsCount(directions, newPosition, fromTo, stepsCount+1, newDirection)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
