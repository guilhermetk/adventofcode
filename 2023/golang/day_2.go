package main

import (
	"fmt"
	"strconv"
	"strings"
)

var games map[int]map[string]int

func exec_day_2() {
	fmt.Println(d2_task1())
}

func d2_task1() int {
	games, err := read_file(2, 1, true)
	sum := 0
	if err == nil {
		for _, game := range games {
			if len(game) > 0 {
				game_split := strings.Split(game, ":")
				game_number, err := strconv.ParseInt(strings.Split(game_split[0], " ")[1], 10, 8)
				game_map := map[string]int{
					"blue":  14,
					"red":   12,
					"green": 13,
				}
				if err == nil {
					cards_split := strings.Split(strings.ReplaceAll(game_split[1], ";", ","), ",")
					should_sum := true
					for _, card := range cards_split {
						card_split := strings.Split(strings.Trim(card, " "), " ")
						quantity, _ := strconv.ParseInt(card_split[0], 10, 8)
						color := card_split[1]
						if quantity > int64(game_map[color]) {
							should_sum = false
						}
					}
					if should_sum {
						sum += int(game_number)
					}
				}
			}
		}
	}
	return sum
}

func d2_task2() {

}
