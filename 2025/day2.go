package main

import (
	"fmt"
	"iter"
	"strconv"
	"strings"
)

func Day2(content []string) {
	input := content[0]
	ranges := strings.SplitSeq(input, ",")
	pt1Result := part1(ranges)
	fmt.Println("part1: ", pt1Result)
	pt2Result := part2(ranges)
	fmt.Println("part2: ", pt2Result)
}

func part2(it iter.Seq[string]) int {
	sum := 0

	for r := range it {
		split := strings.Split(r, "-")
		min, _ := strconv.Atoi(split[0])
		max, _ := strconv.Atoi(split[1])

		for i := min; i <= max; i++ {
			text := strconv.Itoa(i)
		}
	}

	return sum
}

func part1(it iter.Seq[string]) int {
	sum := 0

	for r := range it {
		split := strings.Split(r, "-")
		min, _ := strconv.Atoi(split[0])
		max, _ := strconv.Atoi(split[1])

		for i := min; i <= max; i++ {
			text := strconv.Itoa(i)
			if len(text)%2 == 0 {
				tSplit := len(text) / 2
				p1 := text[0:tSplit]
				p2 := text[tSplit:]

				if p1 == p2 {
					sum += i
				}
				// fmt.Println(text, "_", p1, ",", p2)
			}
		}
	}
	return sum
}
