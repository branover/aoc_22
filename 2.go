package main

import (
	// "strconv"
	// "sort"
	// "fmt"
)

type Move uint8
const (
	Rock Move = 1
	Paper Move = 2
	Scissors Move = 3
)

type Outcome uint8
const (
	Win Outcome = 6
	Lose Outcome = 0
	Draw Outcome = 3
)

var move_map = map[byte]Move{
	'A': Rock,
	'B': Paper,
	'C': Scissors,
	'X': Rock,
	'Y': Paper,
	'Z': Scissors,
}

var outcome_map = map[byte]Outcome{
	'X': Lose,
	'Y': Draw,
	'Z': Win,
}

func calc_winner(left Move, right Move) int {
	if left == right {
		return 3
	}
	if left == Rock && right == Scissors {
		return 0
	}
	if left == Paper && right == Rock {
		return 0
	}
	if left == Scissors && right == Paper {
		return 0
	}
	return 6
}


func calc_move(opponent Move, outcome Outcome) Move {
	if outcome == Win {
		if opponent == Rock {
			return Paper
		}
		if opponent == Paper {
			return Scissors
		}
		if opponent == Scissors {
			return Rock
		}
	}
	if outcome == Draw {
		return opponent
	}
	if outcome == Lose {
		if opponent == Rock {
			return Scissors
		}
		if opponent == Paper {
			return Rock
		}
		if opponent == Scissors {
			return Paper
		}
	}
	return 0
}

func day2a(lines []string) int {
	var solution int	

	for _, line := range lines {
		left := move_map[line[0]]
		right := move_map[line[2]]
		result := calc_winner(left, right)
		solution += result + int(right)	
		// fmt.Printf("%s %d\n", line, result + int(left))	
	}

	return solution
}

func day2b(lines []string) int {
	var solution int	

	for _, line := range lines {
		opponent := move_map[line[0]]
		outcome := outcome_map[line[2]]
		result := calc_move(opponent, outcome)
		solution += int(outcome) + int(result)	
		// fmt.Printf("%s %d %d\n", line, int(outcome), int(result))	
	}

	return solution
}