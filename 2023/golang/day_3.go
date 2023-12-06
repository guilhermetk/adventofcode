package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type number struct {
	value       int
	line_number int
	start_pos   int
	end_pos     int
}

type symbol struct {
	value       string
	line_number int
	position    int
}

func exec_day_3() {
	// fmt.Println(day_1())
	fmt.Println(day_2())
}

func day_1() int {
	sum := 0
	var found_numbers []number
	var found_symbols []symbol
	lines, _ := read_file(3, 1, false)
	for line_index, line := range lines {
		found_number := ""
		start_pos := -1
		for char_position, char := range line {
			is_digit := unicode.IsDigit(char)
			is_period := strings.Contains(string(char), ".")
			if is_digit {
				if found_number == "" {
					start_pos = char_position
				}
				found_number = found_number + string(char)
				if char_position == len(line)-1 {
					number_value, _ := strconv.ParseInt(found_number, 10, 32)
					found_numbers = append(found_numbers, number{int(number_value), line_index, start_pos, char_position - 1})
				}
			} else {
				if found_number != "" {
					number_value, _ := strconv.ParseInt(found_number, 10, 32)
					found_numbers = append(found_numbers, number{int(number_value), line_index, start_pos, char_position - 1})
					found_number = ""
					start_pos = -1
				}
				if !is_period {
					found_symbols = append(found_symbols, symbol{"", line_index, char_position})
				}
			}

		}
	}
	for _, number := range found_numbers {
		should_sum := false
		for _, symbol := range found_symbols {
			is_number_next := is_number_next_to_symbol(number, symbol)
			if is_number_next {
				should_sum = is_number_next
				break
			}
		}
		if should_sum {
			sum += number.value
		}
	}
	return sum
}

func day_2() int {
	sum := 0
	var found_numbers []number
	var found_symbols []symbol
	lines, _ := read_file(3, 1, false)
	for line_index, line := range lines {
		found_number := ""
		start_pos := -1
		for char_position, char := range line {
			is_digit := unicode.IsDigit(char)
			is_period := strings.Contains(string(char), ".")
			if is_digit {
				if found_number == "" {
					start_pos = char_position
				}
				found_number = found_number + string(char)
				if char_position == len(line)-1 {
					number_value, _ := strconv.ParseInt(found_number, 10, 32)
					found_numbers = append(found_numbers, number{int(number_value), line_index, start_pos, char_position - 1})
				}
			} else {
				if found_number != "" {
					number_value, _ := strconv.ParseInt(found_number, 10, 32)
					found_numbers = append(found_numbers, number{int(number_value), line_index, start_pos, char_position - 1})
					found_number = ""
					start_pos = -1
				}
				if !is_period {
					found_symbols = append(found_symbols, symbol{string(char), line_index, char_position})
				}
			}

		}
	}
	fmt.Println(found_symbols)
	for _, symbol := range found_symbols {
		if symbol.value == "*" {
			nearby_numbers := []int{}
			for _, number := range found_numbers {
				is_number_next := is_number_next_to_symbol(number, symbol)
				if is_number_next {
					nearby_numbers = append(nearby_numbers, number.value)
				}
			}
			if len(nearby_numbers) == 2 {
				sum += nearby_numbers[0] * nearby_numbers[1]
			}
		}
	}
	return sum
}

func is_number_next_to_symbol(number number, symbol symbol) bool {
	if abs(symbol.line_number-number.line_number) >= 2 {
		return false
	}

	start_distance := abs(symbol.position - number.start_pos)
	end_distance := abs(symbol.position - number.end_pos)

	if start_distance <= 1 || end_distance <= 1 {
		fmt.Println(number, symbol)
		return true
	}

	return false
}

func abs(value int) int {
	if value < 0 {
		return value * -1
	}
	return value
}
