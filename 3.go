package main

import (
	// "strconv"
	// "sort"
	"strings"
	// "fmt"
)

func get_priority(c int) int {
	if c >= 'A' && c <= 'Z' {
		return c - 38
	}
	return c - 96
}

func day3a(lines []string) int {
	var solution int	

	for _, line := range lines {
		left := line[:len(line)/2]
		right := line[len(line)/2:]
		for _, c := range left {
			if strings.Contains(right, string(c)) {
				priority := get_priority(int(c))
				solution += priority
				break
			}
		}
		
	}

	return solution
}

func day3b(lines []string) int {
	var solution int	

	for i := 0; i < len(lines); i += 3 {
		for _, c := range lines[i] {
			if strings.Contains(lines[i+1], string(c)) && strings.Contains(lines[i+2], string(c)) {
				priority := get_priority(int(c))
				solution += priority
				break
			}
		}
		
	}

	return solution
}