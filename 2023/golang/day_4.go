package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func exec_day_4() {
	fmt.Println(day_4_task_1())
}

func day_4_task_1() int {
	sum := 0
	lines, _ := read_file(4, 1, false)
	for _, line := range lines {
		if len(line) > 0 {
			lines_split := strings.Split(strings.Split(line, ":")[1], "|")
			winning_numbers := stringToIntArray(lines_split[0])
			numbersToCheck := stringToIntArray(lines_split[1])
			count := 0
			for _, win_number := range winning_numbers {
				for _, check_number := range numbersToCheck {
					if win_number == check_number {
						count += 1
					}
				}
			}
			if count > 0 {
				points := 1
				for i := 0; i < count-1; i++ {
					points *= 2
				}
				sum += points
			}
		}
	}
	return sum
}

func stringToIntArray(str string) []int32 {
	var intArray []int32
	split := strings.Split(strings.Trim(str, " "), " ")
	for _, number_str := range split {
		if number_str != "" {
			number_int, _ := strconv.ParseInt(number_str, 10, 32)
			intArray = append(intArray, int32(number_int))

		}
	}
	slices.Sort(intArray)
	return intArray
}
