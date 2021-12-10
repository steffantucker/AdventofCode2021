package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/steffantucker/AdventofCode2021/helpers"
)

func main() {
	testinput := helpers.LoadInputLines("testinput")
	inputinput := helpers.LoadInputLines("input")
	fmt.Printf("test 1: %v\n", puzzle1(testinput))
	fmt.Printf("input 1: %v\n", puzzle1(inputinput))
	fmt.Printf("test 2: %v\n", puzzle2(testinput))
	fmt.Printf("input 2: %v\n", puzzle2(inputinput))
}

func toGrid(in []string) (out [][]int) {
	for _, s := range in {
		r := strings.Split(s, "")
		row := []int{}
		for _, n := range r {
			row = append(row, helpers.MustAtoI(n))
		}
		out = append(out, row)
	}
	return
}

func check(center int, surround []int) bool {
	for _, n := range surround {
		if n <= center {
			return false
		}
	}
	return true
}

func puzzle1(input []string) (out int) {
	grid := toGrid(input)
	for i, row := range grid {
		for j, n := range row {
			sur := []int{}
			if i > 0 {
				sur = append(sur, grid[i-1][j])
			}
			if i < len(grid)-1 {
				sur = append(sur, grid[i+1][j])
			}
			if j > 0 {
				sur = append(sur, grid[i][j-1])
			}
			if j < len(row)-1 {
				sur = append(sur, grid[i][j+1])
			}

			if check(n, sur) {
				out += n + 1
			}
		}
	}
	return
}

type pt struct {
	i, j int
}

func puzzle2(input []string) (out int) {
	grid := toGrid(input)
	bottoms := []pt{}
	for i, row := range grid {
		for j, n := range row {
			sur := []int{}
			if i > 0 {
				sur = append(sur, grid[i-1][j])
			}
			if i < len(grid)-1 {
				sur = append(sur, grid[i+1][j])
			}
			if j > 0 {
				sur = append(sur, grid[i][j-1])
			}
			if j < len(row)-1 {
				sur = append(sur, grid[i][j+1])
			}

			if check(n, sur) {
				bottoms = append(bottoms, pt{i: i, j: j})
			}
		}
	}
	sizes := []int{}
	for _, p := range bottoms {
		points := map[string]int{}
		checkBasin(p.i, p.j, grid, points)
		size := count(points)
		sizes = append(sizes, size)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	return sizes[0] * sizes[1] * sizes[2]
}

func checkBasin(i, j int, grid [][]int, points map[string]int) {
	num := grid[i][j]
	if num == 9 {
		return
	}
	if _, ok := points[fmt.Sprintf("%v,%v", i, j)]; ok {
		return
	}
	points[fmt.Sprintf("%v,%v", i, j)] = num
	height, width := len(grid), len(grid[i])
	left, right, up, down := j-1, j+1, i-1, i+1
	if down < height {
		checkBasin(down, j, grid, points)

	}

	if right < width {
		checkBasin(i, right, grid, points)

	}

	if left >= 0 {
		checkBasin(i, left, grid, points)

	}

	if up >= 0 {
		checkBasin(up, j, grid, points)

	}
}

func count(in map[string]int) (sum int) {

	return len(in)
}

func drawGrid(basinGrid [][]int) {
	for _, row := range basinGrid {
		for _, n := range row {
			fmt.Printf("%v ", n)
		}
		fmt.Println()
	}
}
