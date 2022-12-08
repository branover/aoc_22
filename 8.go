package main

import (
	// "strconv"
	// "sort"
	// "fmt"
)

func build_grid(lines []string) [][]uint8 {
	var grid [][]uint8
	for i, line := range lines {
		grid = append(grid, []uint8{})
		for _, c := range line {
			num := uint8(c - 48)
			grid[i] = append(grid[i], num)
		}
	}
	return grid
}

func day8a(lines []string) int {
	var solution int	
	trees := build_grid(lines)

	for i, row := range trees {
		for j, tree := range row {
			visible_from_left := true
			// Check if all trees to the left are shorter
			for _, left_tree := range row[:j] {
				if left_tree >= tree {
					visible_from_left = false
					break
				}
			}
			visible_from_right := true
			// Check if all trees to the right are shorter
			for _, right_tree := range row[j+1:] {
				if right_tree >= tree {
					visible_from_right = false
					break
				}
			}
			visible_from_above := true
			// Check if all trees above are shorter
			for _, above_tree := range trees[:i] {
				if above_tree[j] >= tree {
					visible_from_above = false
					break
				}
			}
			visible_from_below := true
			// Check if all trees below are shorter
			for _, below_tree := range trees[i+1:] {
				if below_tree[j] >= tree {
					visible_from_below = false
					break
				}
			}
			
			if visible_from_left || visible_from_right || visible_from_above || visible_from_below {
				solution++
				// fmt.Printf("Tree %d at (%d, %d) is visible\n", tree, i, j)
			}			
		}

	}

	return solution
}

func day8b(lines []string) int {
	var solution int	
	var highest_scenic_score int
	trees := build_grid(lines)


	for i, row := range trees {
		for j, tree := range row {
			scenic_score_to_left := 0
			// Check how many trees to left until one is taller
			for k:=j-1; k>=0; k-- {
				scenic_score_to_left++
				if trees[i][k] >= tree {
					break
				}
			}
			scenic_score_to_right := 0
			// Check how many trees to right until one is taller
			for _, right_tree := range row[j+1:] {
				scenic_score_to_right++
				if right_tree >= tree {
					break
				}
			}
			scenic_score_above := 0
			// Check how many trees above until one is taller
			for k:=i-1; k>=0; k-- {
				scenic_score_above++
				if trees[k][j] >= tree {
					break
				}
			}
			scenic_score_below := 0
			// Check how many trees below until one is taller
			for _, below_tree := range trees[i+1:] {
				scenic_score_below++
				if below_tree[j] >= tree {
					break
				}
			}

			scenic_score := scenic_score_to_left * scenic_score_to_right * scenic_score_above * scenic_score_below
			// fmt.Printf("Tree %d at (%d, %d) has scenic score (%d * %d * %d * %d ) = %d\n", tree, i, j, scenic_score_to_left, scenic_score_to_right, scenic_score_above, scenic_score_below , scenic_score)
			if scenic_score > highest_scenic_score {
				highest_scenic_score = scenic_score
			}
		}
	}

	solution = highest_scenic_score
	return solution
}