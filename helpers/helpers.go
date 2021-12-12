package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// LoadInputLines loads the input assuming each line is a string
func LoadInputLines(f string) (lines []string) {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return
}

func LoadNumberGrid(f string, split string) (grid [][]int) {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), split)
		row := []int{}
		for _, n := range line {
			row = append(row, MustAtoI(n))
		}
		grid = append(grid, row)
	}
	return
}

func LoadInputNumbers(f string) (lines []int) {
	l := LoadInputLines(f)
	for _, n := range l {
		num, _ := strconv.Atoi(n)
		lines = append(lines, num)
	}
	return
}

func DrawGrid(grid [][]int) {
	for _, row := range grid {
		for _, n := range row {
			fmt.Printf("%v\t", n)
		}
		fmt.Println()
	}
}

func MustAtoI(in string) int {
	n, _ := strconv.Atoi(in)
	return n
}
