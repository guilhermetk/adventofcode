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

// 0 - UP
// 1 - RIGHT
// 2 - DOWN
// 4 - LEFT

var pipeDirection = map[string][]Position{
	"|": {{x: 0, y: 1}, {x: 0, y: 0}, {x: 0, y: -1}, {x: 0, y: 0}},
	"-": {{x: 0, y: 0}, {x: -1, y: 0}, {x: 0, y: 0}, {x: 1, y: 0}},
	"L": {{x: 1, y: 1}, {x: -1, y: -1}, {x: 0, y: 0}, {x: 0, y: 0}},
	"J": {{x: -1, y: 1}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 1, y: -1}},
	"7": {{x: 0, y: 0}, {x: 0, y: 0}, {x: -1, y: -1}, {x: 1, y: 1}},
	"F": {{x: 0, y: 0}, {x: -1, y: 1}, {x: 1, y: -1}, {x: 0, y: 0}},
	".": {{x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}},
}

var moveDirections = [][]int{
	{0, -1}, //UP
	{1, 0},  //RIGHT
	{0, 1},  //DOWN
	{-1, 0}, //LEFT
}

func main() {
	file, _ := os.ReadFile("./test.txt")
	lines := strings.Split(string(file), "\n")
	var maze [][]string
	var startPosition Position
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
	move(maze, startPosition, startPosition, 0)
}

func move(maze [][]string, startPosition Position, currPos Position, steps int) int {
	// if steps >= 0 && maze[startPosition.x][startPosition.y] == maze[currPos.x][currPos.y] {
	// 	return steps
	// }
	for i := 0; i < 4; i++ {
		newPipe := maze[currPos.x+moveDirections[i][1]][currPos.y+moveDirections[i][0]]

	}
	return 0
}

func task2() {
}
