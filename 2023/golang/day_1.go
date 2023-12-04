package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	fmt.Fprintln("Task 1 result is: %d", task1())
}

func task1() int {
	lines, err := read_file(1, false)
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
			fmt.Println(line_number)
			value, err := strconv.ParseInt(line_number, 10, 8)
			if err == nil {
				sum += int(value)
			}
		}
	}
	return sum
}
