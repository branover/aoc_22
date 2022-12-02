package main

import (
	"strconv"
	"sort"
)

func day1a(lines []string) int {
	var solution int	
	var elves []int
	var elf int
	for _, line := range lines {
		cals, err := strconv.Atoi(line)
		if err != nil {
			elves = append(elves, elf)
			elf = 0
			continue
		}
		elf += cals
		
	}

	for _, elf := range elves {
		if elf > solution {
			solution = elf
		}
	}


	return solution
}

func day1b(lines []string) int {
	var solution int	
	var elves []int
	var elf int
	for _, line := range lines {
		cals, err := strconv.Atoi(line)
		if err != nil {
			elves = append(elves, elf)
			elf = 0
			continue
		}
		elf += cals
		
	}

	sort.Slice(elves, func(i, j int) bool { return elves[i] < elves[j] })
	top3 := elves[len(elves)-3:]
	for _, elf := range top3 {
		solution += elf
	}

	return solution
}