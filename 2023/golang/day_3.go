package main

import (
	"fmt"
)

func exec_day_3() {
	day_1()
}

func day_1() {
	lines, _ := read_file(3, 1, false)
	for _, line := range lines {
		for i := 0; i < len(line); i++ {
			// is_digit := unicode.IsDigit(rune(line[i]))
			// is_period := strings.Contains(string(line[i]), ".")
			// if is_digit && !is_period {
			// fmt.Print(line[i])
			// } else {
			// fmt.Print(" ")
		}
		fmt.Println(line)
	}
}
