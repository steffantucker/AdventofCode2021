package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/steffantucker/AdventofCode2021/helpers"
)

type bingoBoard struct {
	nums    [][]string
	stamped map[string]bool
}

func (b bingoBoard) isBingo() bool {
	for i, row := range b.nums {
		rcount := 0
		ccount := 0
		for j := range row {
			rhash := fmt.Sprintf("%v%v", i, j)
			chash := fmt.Sprintf("%v%v", j, i)
			if _, ok := b.stamped[rhash]; ok {
				rcount++
			}
			if _, ok := b.stamped[chash]; ok {
				ccount++
			}
		}
		if rcount == len(row) || ccount == len(row) {
			return true
		}
	}
	return false
}

func (b bingoBoard) stamp(s string) bool {
	for i, row := range b.nums {
		for j, num := range row {
			if s == num {
				shash := fmt.Sprintf("%v%v", i, j)
				b.stamped[shash] = true
				return b.isBingo()
			}
		}
	}
	return false
}

func (b bingoBoard) score(in string) int {
	score := 0
	for i, row := range b.nums {
		for j, n := range row {
			hash := fmt.Sprintf("%v%v", i, j)
			if _, ok := b.stamped[hash]; !ok {
				num, _ := strconv.Atoi(n)
				score += num
			}
		}
	}
	num, _ := strconv.Atoi(in)
	return score * num
}

func (b bingoBoard) draw() (out string) {
	for i, row := range b.nums {
		rowString := []string{}
		for j, n := range row {
			hash := fmt.Sprintf("%v%v", i, j)
			if _, ok := b.stamped[hash]; ok {
				rowString = append(rowString, "x")
			} else {
				rowString = append(rowString, n)
			}
		}
		out += fmt.Sprintf("%v\n", strings.Join(rowString, "\t"))
	}
	return
}

func newBoard(b []string) bingoBoard {
	board := [][]string{}
	for _, brow := range b {
		board = append(board, strings.Split(brow, " "))
	}
	stamped := make(map[string]bool, 4)
	return bingoBoard{
		nums:    board,
		stamped: stamped,
	}
}

func main() {
	testinput := helpers.LoadInputLines("testinput")
	inputinput := helpers.LoadInputLines("input")
	fmt.Printf("test 1: %v\ninput 1: %v\n", puzzle1(testinput), puzzle1(inputinput))
	fmt.Printf("test 2: %v\ninput: %v\n", puzzle2(testinput), puzzle2(inputinput))
}

func buildNumsandBoards(input []string) ([]string, []bingoBoard) {

	nums := strings.Split(input[0], ",")
	boards := []bingoBoard{}
	b := []string{}
	for _, row := range input[2:] {
		if row == "" {
			boards = append(boards, newBoard(b))
			b = []string{}
			continue
		}
		b = append(b, row)
	}
	boards = append(boards, newBoard(b))
	return nums, boards
}

func puzzle1(input []string) (out int) {
	nums, boards := buildNumsandBoards(input)
	for _, num := range nums {
		for _, board := range boards {
			if board.stamp(num) {
				fmt.Printf("num:%v\nboard:", num)
				fmt.Println(board.draw())
				return board.score(num)
			}
		}
	}
	return -1
}

func puzzle2(input []string) int {
	nums, boards := buildNumsandBoards(input)
	fmt.Println(nums)
	for _, num := range nums {
		for i, l := 0, len(boards); i < l; i++ {
			board := boards[i]
			if board.stamp(num) {

				if len(boards) == 1 {
					return board.score(num)
				}
				boards[i] = boards[len(boards)-1]
				boards = boards[:len(boards)-1]
				fmt.Printf("num:%v\nboard:\n", num)
				fmt.Println(board.draw())
				i--
				l--
			}
		}
	}
	return -1
}
