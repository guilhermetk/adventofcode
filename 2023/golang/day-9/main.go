package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("prod.txt")
	lines := strings.Split(string(file), "\n")
	historyData := [][]int{}
	for _, line := range lines {
		if line != "" {
			lineSplit := strings.Split(line, " ")
			historyData = append(historyData, stringArrayToIntArray(lineSplit))
		}
	}
	task1(historyData)
}

func task1(data [][]int) {
	sum := 0
	for _, item := range data {
		value := sequenceFromDiff(item)
		sum += value
	}
	fmt.Println(sum)
}

func sequenceFromDiff(data []int) int {
	newSequence := []int{}
	for i := 0; i < len(data)-1; i++ {
		numToAppend := data[i+1] - data[i]
		newSequence = append(newSequence, numToAppend)
	}

	allZeros := true
	for _, item := range data {
		if item != 0 {
			allZeros = false
		}
	}
	if !allZeros {
		response := sequenceFromDiff(newSequence)
		return data[0] - response
	}

	return 0
}

func task2() {
}

func stringArrayToIntArray(arr []string) []int {
	number := []int{}
	for _, item := range arr {
		intValue, _ := strconv.Atoi(item)
		number = append(number, intValue)
	}
	return number
}
