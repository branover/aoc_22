package main

import (
	"strconv"
	// "sort"
	"strings"
	// "fmt"
)

type num_range struct {
	start int
	end int
}

func build_num_range(line string) num_range {
	s := strings.SplitN(line, "-", 2)

	start_s, end_s := s[0], s[1]
	start, _ := strconv.Atoi(start_s)
	end, _ := strconv.Atoi(end_s)
	return num_range{start, end}
}

func get_ranges(line string) (num_range, num_range) {

	s := strings.SplitN(line, ",", 2)
	left, right := s[0], s[1]

	return build_num_range(left), build_num_range(right)
}

func range_contains(range1, range2 num_range) bool {
	if range1.start <= range2.start && range1.end >= range2.end {
		return true
	}
	if range2.start <= range1.start && range2.end >= range1.end {
		return true
	}
	return false
}

func range_overlaps(range1, range2 num_range) bool {
	if range1.start <= range2.start && range1.end >= range2.start {
		return true
	}
	if range2.start <= range1.start && range2.end >= range1.start {
		return true
	}
	return false
}

func day4a(lines []string) int {
	var solution int	

	for _, line := range lines {
		left, right := get_ranges(line)
		if range_contains(left, right) {
			solution++
		}		
	}

	return solution
}

func day4b(lines []string) int {
	var solution int

	for _, line := range lines {
		left, right := get_ranges(line)
		if range_overlaps(left, right) {
			solution++
		}		
	}

	return solution
}