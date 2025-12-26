package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, _ := readFile(false)
	// Day1(content)
	// Day2(content)
	Day3(content)
}

func readFile(testInput bool) ([]string, error) {
	var fileType string
	if testInput {
		fileType = "test"
	} else {
		fileType = "result"
	}
	fileName := fmt.Sprintf("./input_%s.txt", fileType)
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}
