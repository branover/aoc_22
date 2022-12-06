package main

import (
	// "strconv"
	// "sort"
	"strings"
	"fmt"
	"strconv"
)

func day5a(lines []string) int {
	var solution int	
	var move_start int
	var stacks [][]byte
 

	for i, line := range lines {
		if strings.Contains(line, "1") {
			move_start = i+1
			break
		}
	}

	for i := move_start; i>= 0; i-- {
		line := lines[i]
		for n := 0; n < len(line); n+=4 {
			stack_num := n/4
			if stack_num >= len(stacks) {
				stacks = append(stacks, []byte{})
			}
			box := line[n:n+3]
			if box[0] == '[' {
				stacks[stack_num] = append(stacks[stack_num], box[1])
			}
			// fmt.Println(line[n:n+3])

		}		
	}

	for i:= move_start+1; i < len(lines); i++ {
		line := strings.SplitN(lines[i], " ", 6)
		num, from, to := line[1], line[3], line[5]
		num_i, _ := strconv.Atoi(num)
		from_i, _ := strconv.Atoi(from)
		to_i, _ := strconv.Atoi(to)
		from_i -= 1
		to_i -= 1

		for j:= 0; j < num_i; j++ {
			stacks[to_i] = append(stacks[to_i], stacks[from_i][len(stacks[from_i]) - 1])
			stacks[from_i] = stacks[from_i][:len(stacks[from_i])-1]
		}
	}


	var solution_s string
	for _, stack := range stacks {
		solution_s += string(stack[len(stack)-1])
	}	
	fmt.Println(solution_s)

	return solution
}

func day5b(lines []string) int {
	var solution int	
	var move_start int
	var stacks [][]byte
 

	for i, line := range lines {
		if strings.Contains(line, "1") {
			move_start = i+1
			break
		}
	}

	for i := move_start; i>= 0; i-- {
		line := lines[i]
		for n := 0; n < len(line); n+=4 {
			stack_num := n/4
			if stack_num >= len(stacks) {
				stacks = append(stacks, []byte{})
			}
			box := line[n:n+3]
			if box[0] == '[' {
				stacks[stack_num] = append(stacks[stack_num], box[1])
			}
		}		
	}

	for i:= move_start+1; i < len(lines); i++ {
		line := strings.SplitN(lines[i], " ", 6)
		num, from, to := line[1], line[3], line[5]
		num_i, _ := strconv.Atoi(num)
		from_i, _ := strconv.Atoi(from)
		to_i, _ := strconv.Atoi(to)
		from_i -= 1
		to_i -= 1

		for j:= 0; j< num_i; j++ {
			to_move := len(stacks[from_i]) - num_i + j
			stacks[to_i] = append(stacks[to_i], stacks[from_i][to_move])
		}
		stacks[from_i] = stacks[from_i][:len(stacks[from_i]) - num_i]

	}


	var solution_s string
	for _, stack := range stacks {
		solution_s += string(stack[len(stack)-1])
	}	
	fmt.Println(solution_s)

	return solution
}