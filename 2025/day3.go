package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/guilhermetk/adventofcode/stack"
)

func Day3(content []string) {
	pt1Result := part1d3(content)
	fmt.Println("part1: ", pt1Result)
	pt2Result := part2d3(content)
	fmt.Println("part2: ", pt2Result)
}

func part1d3(content []string) int {
	sum := 0

	for _, bank := range content {
		largest := 0
		for j := 0; j < len(bank); j++ {
			for k := j + 1; k < len(bank); k++ {
				number, _ := strconv.Atoi(string([]byte{bank[j], bank[k]}))

				if number > largest {
					largest = number
				}
			}
		}
		sum += largest
	}

	return sum
}

func part2d3(content []string) int {
	sum := 0

	for _, bank := range content {
		k := len(bank) - 12
		largest := remomeKDigits(bank, k)
		fmt.Println(bank, ": ", largest)
		sum += largest
	}

	return sum
}

func remomeKDigits(s string, k int) int {
	if k >= len(s) {
		return 0
	}

	digits := strings.Split(s, "")
	stack := stack.New[int]()

	for _, sDigit := range digits {
		digit, _ := strconv.Atoi(sDigit)
		for k > 0 && stack.Len() > 0 && stack.Peek() < int(digit) {
			stack.Pop()
			k--
		}

		stack.Add(digit)
	}

	for k > 0 && stack.Len() > 0 {
		stack.Pop()
		k--
	}

	value := stack.String()
	intValue, _ := strconv.Atoi(value)

	return intValue
}
