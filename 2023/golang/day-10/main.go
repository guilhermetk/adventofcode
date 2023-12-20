package main

import (
	"fmt"
	"os"
	"strings"
)

type Position struct {
	x int
	y int
}

var visitedPositions map[Position]bool

// COMING FROM:
// 0 - UP
// 1 - RIGHT
// 2 - DOWN
// 4 - LEFT
var pipeDirection = map[string][]bool{
	"|": {true, false, true, false},
	"-": {false, true, false, true},
	"L": {false, false, true, true},
	"J": {false, true, true, false},
	"7": {true, true, false, false},
	"F": {true, false, false, true},
	".": {false, false, false, false},
	"S": {false, false, false, false},
}

var moveDirections = [][]int{
	{-1, 0}, //UP
	{0, 1},  //RIGHT
	{1, 0},  //DOWN
	{0, -1}, //LEFT
}

var path []string = []string{}

func main() {
	file, _ := os.ReadFile("./test.txt")
	lines := strings.Split(string(file), "\n")
	var maze [][]string
	var startPosition Position
	visitedPositions = make(map[Position]bool)
	fmt.Println(visitedPositions[Position{x: 1, y: 1}])
	for index, line := range lines {
		if len(line) > 0 {
			maze = append(maze, strings.Split(line, ""))
			startIndex := strings.IndexRune(line, 'S')
			if startIndex >= 0 {
				startPosition = Position{x: startIndex, y: index}
			}
		}
	}
	fmt.Println(startPosition)
	task1(maze, startPosition)
}

func task1(maze [][]string, startPosition Position) {
	fmt.Println(move(maze, startPosition, startPosition, 0))
}

func move(maze [][]string, startPosition Position, currPos Position, steps int) int {
	visitedPositions[currPos] = true
	currentPosPipe := maze[currPos.x][currPos.y]
	if steps > 0 && currentPosPipe == "S" {
		fmt.Println("Encerrou: ", steps)
		return steps
	}

	maxSteps := steps

	for i := 0; i < 4; i++ {
		moveX := moveDirections[i][0]
		moveY := moveDirections[i][1]
		newPipePosition := Position{x: currPos.x + moveX, y: currPos.y + moveY}
		newPipe := maze[currPos.x+moveX][currPos.y+moveY]
		canMoveToPipe := pipeDirection[newPipe][i]
		if !outOfBounds(maze, newPipePosition) && canMoveToPipe && !visitedPositions[newPipePosition] {
			fmt.Println(newPipe)
			newStep := move(maze, startPosition, newPipePosition, maxSteps+1)
			if newStep > maxSteps {
				maxSteps = newStep
			}
		}
	}
	return maxSteps
}

func outOfBounds(maze [][]string, pos Position) bool {
	if pos.x < 0 || pos.y < 0 || pos.x > len(maze[0])-1 || pos.y > len(maze)-1 {
		return true
	}
	return false
}

func canMoveToPipe(pos Position) bool {
	if pos.x == 0 && pos.y == 0 {
		return false
	}
	return true
}

func task2() {
}
