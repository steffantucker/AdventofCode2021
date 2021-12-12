package main

import (
	"fmt"

	"github.com/steffantucker/AdventofCode2021/helpers"
)

func main() {
	testinput := helpers.LoadNumberGrid("testinput", "")
	inputinput := helpers.LoadNumberGrid("input", "")
	fmt.Printf("test 1: %v\n", puzzle1(testinput, 100))
	//fmt.Printf("input 1: %v\n", puzzle1(inputinput, 100))
	fmt.Printf("test 2: %v\n", puzzle2(testinput))
	fmt.Printf("input 2: %v\n", puzzle2(inputinput))
}

func buildUp(grid [][]int) [][]int {
	for i, row := range grid {
		for j := range row {
			grid[i][j]++
		}
	}
	return grid
}

func sparkle(i, j int, grid [][]int) (int, [][]int) {
	count := 0
	grid[i][j]++
	if grid[i][j] != 11 && grid[i][j] != 10 {
		return count, grid
	}
	count++
	grid[i][j] = 12
	for ii := i - 1; ii <= i+1; ii++ {
		if ii < 0 || ii >= len(grid) {
			continue
		}
		for jj := j - 1; jj <= j+1; jj++ {
			if jj < 0 || jj >= len(grid[0]) || (ii == i && jj == j) {
				continue
			}
			c, g := sparkle(ii, jj, grid)
			count += c
			grid = g
		}
	}
	return count, grid
}

func calmDown(grid [][]int) [][]int {
	for i, row := range grid {
		for j, o := range row {
			if o > 9 {
				grid[i][j] = 0
			}
			if o == 10 {
				helpers.DrawGrid(grid)
				panic("not everyone flashed!")
			}
		}
	}
	return grid
}

func puzzle1(input [][]int, steps int) (out int) {
	grid := input
	count := 0
	for step := 0; step < steps; step++ {
		grid = buildUp(grid)
		for i, row := range grid {
			for j, o := range row {
				if o == 10 {
					count, grid = sparkle(i, j, grid)
					out += count
				}
			}
		}
		grid = calmDown(grid)
		fmt.Println(step)
		helpers.DrawGrid(grid)
	}
	return
}

func puzzle2(input [][]int) (out int) {
	grid := input
	count := 0
	helpers.DrawGrid(grid)
	fmt.Println()
	for ; !simulFlash(grid); count++ {
		grid = buildUp(grid)
		for i, row := range grid {
			for j, o := range row {
				if o == 10 {
					_, grid = sparkle(i, j, grid)
				}
			}
		}
		grid = calmDown(grid)
		if simulFlash(grid) {
			break
		}
	}
	helpers.DrawGrid(grid)
	return count + 1
}

func simulFlash(grid [][]int) bool {
	for _, row := range grid {
		for _, n := range row {
			if n != 0 {
				return false
			}
		}
	}
	return true
}
