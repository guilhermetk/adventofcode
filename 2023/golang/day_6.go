package main

import (
	"fmt"
	"strings"
)

func exec_day_6() {
	// fmt.Println(d6_task1())
	fmt.Println(d6_task2())
}

func d6_task1() int {
	lines, _ := read_file(6, 1, false)
	raceTimes := stringToUnsortedIntArray(strings.Split(lines[0], ":")[1])
	raceRecords := stringToUnsortedIntArray(strings.Split(lines[1], ":")[1])
	fmt.Println(raceTimes)
	fmt.Println(raceRecords)
	mult := 0
	for i := 0; i < len(raceTimes); i++ {
		raceTime := raceTimes[i]
		count := 0
		for j := 0; j <= int(raceTime); j++ {
			raceDistance := calculateDistante(j, int(raceTime))
			if raceDistance > int(raceRecords[i]) {
				count += 1
			}
		}
		if mult == 0 {
			mult = count
		} else {
			mult *= count
		}
	}
	return mult
}

func d6_task2() int {
	lines, _ := read_file(6, 1, false)
	raceTimes := stringToUnsortedIntArray(strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", ""))
	raceRecords := stringToUnsortedIntArray(strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", ""))
	fmt.Println(raceTimes)
	fmt.Println(raceRecords)
	mult := 0
	for i := 0; i < len(raceTimes); i++ {
		raceTime := raceTimes[i]
		count := 0
		for j := 0; j <= int(raceTime); j++ {
			raceDistance := calculateDistante(j, int(raceTime))
			if raceDistance > int(raceRecords[i]) {
				count += 1
			}
		}
		if mult == 0 {
			mult = count
		} else {
			mult *= count
		}
	}
	return mult
}

func calculateDistante(buttonHold int, raceTime int) int {
	return (raceTime - buttonHold) * buttonHold
}
