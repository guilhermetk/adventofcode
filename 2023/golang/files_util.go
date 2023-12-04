package main

import (
	"fmt"
	"os"
	"strings"
)

func read_file(day int8, use_test_input bool) ([]string, error) {
	var file_type string
	if use_test_input {
		file_type = "test"
	} else {
		file_type = "prod"
	}
	file_name := fmt.Sprintf("./input/day_%d_%s.txt", day, file_type)
	data, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return nil, err
	}
	return strings.Split(string(data), "\n"), nil
}
