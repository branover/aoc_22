package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func readlines(path string) []string {
	var lines []string
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func main() {
	args := os.Args[1:]
	day := args[0]
	part := args[1]

	lines := readlines("data/" + day)
	// lines := readlines("sample/" + day)

	var func_call func([]string) int

	switch day + part {
	case "1a":
		func_call = day1a
	case "1b":
		func_call = day1b
	case "2a":
		func_call = day2a
	case "2b":
		func_call = day2b
	case "3a":
		func_call = day3a
	case "3b":
		func_call = day3b
	case "4a":
		func_call = day4a
	case "4b":
		func_call = day4b
	case "5a":
		func_call = day5a
	case "5b":
		func_call = day5b
	case "6a":
		func_call = day6a
	case "6b":
		func_call = day6b
	case "7a":
		func_call = day7a
	case "7b":
		func_call = day7b
	case "8a":
		func_call = day8a
	case "8b":
		func_call = day8b
	case "9a":
		func_call = day9a
	case "9b":
		func_call = day9b
	case "10a":
		func_call = day10a
	case "10b":
		func_call = day10b
	case "11a":
		func_call = day11a
	case "11b":
		func_call = day11b
	case "12a":
		func_call = day12a
	case "12b":
		func_call = day12b
	case "13a":
		func_call = day13a
	case "13b":
		func_call = day13b
	case "14a":
		func_call = day14a
	case "14b":
		func_call = day14b
	case "15a":
		func_call = day15a
	case "15b":
		func_call = day15b
	case "16a":
		func_call = day16a
	case "16b":
		func_call = day16b
	case "17a":
		func_call = day17a
	case "17b":
		func_call = day17b
	case "18a":
		func_call = day18a
	case "18b":
		func_call = day18b
	case "19a":
		func_call = day19a
	case "19b":
		func_call = day19b
	case "20a":
		func_call = day20a
	case "20b":
		func_call = day20b
	case "21a":
		func_call = day21a
	case "21b":
		func_call = day21b
	case "22a":
		func_call = day22a
	case "22b":
		func_call = day22b
	case "23a":
		func_call = day23a
	case "23b":
		func_call = day23b
	case "24a":
		func_call = day24a
	case "24b":
		func_call = day24b
	case "25a":
		func_call = day25a
	case "25b":
		func_call = day25b
	}

	max_repetitions := 1000
	var total_elapsed time.Duration
	for i := 0; i < max_repetitions; i++ {
		start := time.Now()
		func_call(lines)
		total_elapsed += time.Since(start)
	}
	start := time.Now()
	solution := func_call(lines)
	elapsed := time.Since(start)

	fmt.Println("Solution: ", solution)
	fmt.Println("Time elapsed: ", elapsed.Microseconds(), "µs")
	fmt.Println("Average time elapsed: ", total_elapsed.Microseconds()/int64(max_repetitions), "µs")

}
