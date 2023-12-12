package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func exec_day_5() {
	task2()
}

func task1() {
	blocks, _ := read_file_by_space_blocks(5, 1, false)
	seeds := []int64{}
	commands := [][]int64{}
	for blockNumber, block := range blocks {
		if blockNumber == 0 {
			seeds = stringToUnsortedIntArray(strings.Split(block, ":")[1])
		} else {
			commands_string := strings.Split(block, "\n")[1:]
			mapCommands := []int64{}
			for _, command_string := range commands_string {
				if command_string != "" {
					mapCommands = append(mapCommands, stringToUnsortedIntArray(command_string)...)
				}
			}
			commands = append(commands, mapCommands)
		}
	}
	var lowest int64 = math.MaxInt64
	fmt.Println(commands)
	for _, seed := range seeds {
		destination := parseSourceToDestination(seed, commands, 0, 0)
		fmt.Println(seed, destination)
		if destination < lowest {
			lowest = destination
		}
	}
	fmt.Println(lowest)
}

func task2() {
	blocks, _ := read_file_by_space_blocks(5, 1, false)
	seeds := []int64{}
	commands := [][]int64{}
	for blockNumber, block := range blocks {
		if blockNumber == 0 {
			seeds = stringToUnsortedIntArray(strings.Split(block, ":")[1])
		} else {
			commands_string := strings.Split(block, "\n")[1:]
			mapCommands := []int64{}
			for _, command_string := range commands_string {
				if command_string != "" {
					mapCommands = append(mapCommands, stringToUnsortedIntArray(command_string)...)
				}
			}
			commands = append(commands, mapCommands)
		}
	}
	var lowest int64 = math.MaxInt64

	for i := 0; i+1 < len(seeds); i += 2 {
		seed := seeds[i]
		seedRange := seeds[i+1]

		for j := seed; j < seed+seedRange; j++ {
			go goRoutine(j, commands)
		}
	}

	fmt.Println(lowest)
}

var lowest int64 = math.MaxInt64

func goRoutine(j int64, commands [][]int64) {
	destination := parseSourceToDestination(j, commands, 0, 0)
	if destination < lowest {
		lowest = destination
		fmt.Println(lowest)
	}
}

func parseSourceToDestination(source int64, commands [][]int64, mapIndex int, commandIndex int) int64 {
	if mapIndex > len(commands)-1 {
		return source
	}

	destinationRangeStart := commands[mapIndex][commandIndex+0]
	sourceRangeStart := commands[mapIndex][commandIndex+1]
	sourceRange := commands[mapIndex][commandIndex+2]
	difference := sourceRangeStart - destinationRangeStart
	parsedSource := source

	if source >= sourceRangeStart && source < sourceRangeStart+sourceRange {
		parsedSource = source - difference
		return parseSourceToDestination(parsedSource, commands, mapIndex+1, 0)
	}

	if commandIndex+3 >= len(commands[mapIndex]) {
		return parseSourceToDestination(parsedSource, commands, mapIndex+1, 0)
	}

	return parseSourceToDestination(parsedSource, commands, mapIndex, commandIndex+3)
}

func stringToUnsortedIntArray(str string) []int64 {
	var intArray []int64
	split := strings.Split(strings.Trim(str, " "), " ")
	for _, number_str := range split {
		if number_str != "" {
			number_int, _ := strconv.ParseInt(number_str, 10, 64)
			intArray = append(intArray, int64(number_int))

		}
	}
	return intArray
}
