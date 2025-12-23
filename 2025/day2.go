package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Day2(content []string) {
	input := content[0]
	ranges := slices.Collect(strings.SplitSeq(input, ","))
	// pt1Result := part1(ranges)
	// fmt.Println("part1: ", pt1Result)
	pt2Result := part2(ranges)
	fmt.Println("part2: ", pt2Result)
}

func part2(it []string) int {
	sum := 0

	for _, r := range it {
		split := strings.Split(r, "-")
		min, _ := strconv.Atoi(split[0])
		max, _ := strconv.Atoi(split[1])

		for i := min; i <= max; i++ {
			text := strconv.Itoa(i)

			for j := 1; j <= len(text)/2; j++ {
				prefix := text[:j]
				rest := text[j:]

				result := strings.ReplaceAll(rest, prefix, "")
				if result == "" {
					sum += i
					fmt.Println(text, "_", sum)
					break
				}
			}
		}
	}

	return sum
}

func part1(it []string) int {
	sum := 0

	for _, r := range it {
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
