package main

import (
	"strconv"
	// "sort"
)

func day1a(lines []string) int {
	var solution int	
	var elf int
	for _, line := range lines {
		cals, err := strconv.Atoi(line)
		if err != nil {
			elf = 0
			continue
		}
		elf += cals
		if elf > solution {
			solution = elf
		}
	}

	return solution
}

func day1b(lines []string) int {
	var solution int	
	var top3 [3]int
	var elf int
	for _, line := range lines {
		cals, err := strconv.Atoi(line)
		if err != nil {
			elf = 0
			continue
		}
		elf += cals

		for i, top := range top3 {
			if elf > top {
				top3[i] = elf
				break
			}
		}
	}

	for _, elf := range top3 {
		solution += elf
	}

	return solution
}