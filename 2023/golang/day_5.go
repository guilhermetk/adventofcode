package main

import (
	"fmt"
)

func exec_day_5() {
	task1()
}

func task1() {
	blocks, _ := read_file_by_space_blocks(5, 1, true)
	// seeds := []int{}
	for number, block := range blocks {
		fmt.Println(number, block)
	}
}

func calculateDestination(seed int) {

}
