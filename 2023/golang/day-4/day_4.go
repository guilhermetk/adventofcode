package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func exec_day_4() {
	fmt.Println(day_4_task_2())
}

func day_4_task_1() int {
	var sum int = 0
	lines, _ := read_file(4, 1, false)
	for _, line := range lines {
		if len(line) > 0 {
			lines_split := strings.Split(strings.Split(line, ":")[1], "|")
			winning_numbers := stringToIntArray(lines_split[0])
			numbersToCheck := stringToIntArray(lines_split[1])
			matched_numbers := match_winning_numbers(winning_numbers, numbersToCheck)
			sum += calculate_points(matched_numbers)
		}
	}
	return sum
}

func day_4_task_2() int {
	count := 0
	lines, _ := read_file(4, 1, false)
	matches_by_card := make(map[int][]int)
	max_index := 0
	for line_index, line := range lines {
		if len(line) > 0 {
			lines_split := strings.Split(strings.Split(line, ":")[1], "|")
			winning_numbers := stringToIntArray(lines_split[0])
			numbersToCheck := stringToIntArray(lines_split[1])
			matched_numbers_count := match_winning_numbers(winning_numbers, numbersToCheck)
			matched_cards := []int{}
			for i := line_index + 1; i <= line_index+matched_numbers_count; i++ {
				matched_cards = append(matched_cards, i+1)
			}
			matches_by_card[line_index+1] = matched_cards
			max_index = matched_numbers_count + 1
		}
		if line_index == max_index {
			break
		}
	}
	for i := 1; i <= len(matches_by_card); i++ {
		count += dive(matches_by_card, i) + 1
	}
	return count
}

func dive(matches map[int][]int, card int) int {
	if len(matches[card]) == 0 {
		return 0
	}
	copies := matches[card]
	sum := len(copies)
	for _, copie := range copies {
		sum += dive(matches, copie)
	}
	return sum
}

func match_winning_numbers(winning_numbers []int, numbersToCheck []int) int {
	count := 0
	for _, win_number := range winning_numbers {
		for _, check_number := range numbersToCheck {
			if win_number == check_number {
				count += 1
			}
		}
	}
	return count
}

func calculate_points(count int) int {
	sum := 0
	if count > 0 {
		points := 1
		for i := 0; i < count-1; i++ {
			points *= 2
		}
		sum += points
	}
	return sum
}

func stringToIntArray(str string) []int {
	var intArray []int
	split := strings.Split(strings.Trim(str, " "), " ")
	for _, number_str := range split {
		if number_str != "" {
			number_int, _ := strconv.ParseInt(number_str, 10, 32)
			intArray = append(intArray, int(number_int))

		}
	}
	slices.Sort(intArray)
	return intArray
}
