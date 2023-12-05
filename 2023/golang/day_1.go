package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var digit_map map[string]int

func exec_day_1() {
	digit_map = make(map[string]int)
	digit_map["one"] = 1
	digit_map["two"] = 2
	digit_map["three"] = 3
	digit_map["four"] = 4
	digit_map["five"] = 5
	digit_map["six"] = 6
	digit_map["seven"] = 7
	digit_map["eight"] = 8
	digit_map["nine"] = 9
	fmt.Println(d1_task1())
	fmt.Println(d1_task2())
}

func d1_task1() int {
	lines, err := read_file(1, 1, false)
	sum := 0
	if err == nil {
		for _, line := range lines {
			var first string
			var last string
			for i := 0; i < len(line); i++ {
				var j = len(line) - i - 1
				if first == "" && unicode.IsDigit(rune(line[i])) {
					first = string(line[i])
				}
				if last == "" && unicode.IsDigit(rune(line[j])) {
					last = string(line[j])
				}
				if first != "" && last != "" {
					break
				}
			}
			line_number := first + last
			value, err := strconv.ParseInt(line_number, 10, 8)
			if err == nil {
				sum += int(value)
			}
		}
	}
	return sum
}

func d1_task2() int {
	lines, err := read_file(1, 2, false)
	sum := 0
	if err == nil {
		for _, line := range lines {
			var first string
			var last string
			for i := 0; i < len(line); i++ {
				var j = len(line) - i - 1
				if first == "" && unicode.IsDigit(rune(line[i])) {
					first = string(line[i])
				}

				if first == "" {
					first_substr := line[0 : i+1]
					matched_digit := match_digit(first_substr)
					if matched_digit > 0 {
						first = fmt.Sprintf("%d", matched_digit)
					}
				}

				if last == "" && unicode.IsDigit(rune(line[j])) {
					last = string(line[j])
				}

				if last == "" {
					last_substr := line[len(line)-i-1:]
					matched_digit := match_digit(last_substr)
					if matched_digit > 0 {
						last = fmt.Sprintf("%d", matched_digit)
					}
				}

				if first != "" && last != "" {
					break
				}
			}
			line_number := first + last
			value, err := strconv.ParseInt(line_number, 10, 8)
			if err == nil {
				sum += int(value)
			}
		}
	}
	return sum
}

func match_digit(str string) int {
	for key, value := range digit_map {
		if strings.Contains(str, key) {
			return value
		}
	}
	return 0
}
