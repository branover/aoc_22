package main

import (
	// "strconv"
	// "fmt"
)

func find_sequence(line string, seq_length int) int {
	length := len(line)
	for i := 0; i < len(line); i++ {
		repeat := false

		for j := 1; j <= seq_length - 1; j++ {
			if i < j {
				break
			}
			// fmt.Printf("i:%d j:%d %c %s\n", i, j, line[i], line)

			if i >= j && line[i] == line [i-j] {
				line = line[i - j + 1:]
				i = -1
				// fmt.Println("repeat")
				repeat = true
				break
			}
		} 
		if !repeat && i >= (seq_length - 1) {
			return length - len(line) + seq_length
		}
	}
	return 0
				
}

func day6a(lines []string) int {
	var solution int	
	seq_length := 4

	for _, line := range lines {
		solution = find_sequence(line, seq_length)
		// fmt.Println(solution)				
	}

	return solution
}

func day6b(lines []string) int {
	var solution int	
	seq_length := 14

	for _, line := range lines {
		solution = find_sequence(line, seq_length)
		// fmt.Println(solution)				
	}

	return solution
}