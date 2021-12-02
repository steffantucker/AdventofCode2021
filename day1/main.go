package main

import (
	"fmt"

	"github.com/steffantucker/AdventofCode2021/helpers"
)

func main() {
	testinput := helpers.LoadInputNumbers("testinput")
	inputinput := helpers.LoadInputNumbers("input")
	fmt.Printf("test 1: %v\ninput 1: %v\ntest 2: %v\ninput: %v\n", puzzle1(testinput), puzzle1(inputinput), puzzle2(testinput), puzzle2(inputinput))
}

func puzzle1(input []int) (out int) {
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			out++
		}
	}
	return
}

func puzzle2(input []int) int {
	in := []int{}
	for i := 2; i < len(input); i++ {
		in = append(in, input[i]+input[i-1]+input[i-2])
	}
	return puzzle1(in)
}
