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
			if elf > solution {
				solution = elf
			}			
			elf = 0
			continue
		}
		elf += cals

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
			for i, top := range top3 {
				if elf > top {
					top3[i] = elf
					break
				}
			}			
			elf = 0
			continue
		}
		elf += cals


	}

	for _, elf := range top3 {
		solution += elf
	}

	return solution
}