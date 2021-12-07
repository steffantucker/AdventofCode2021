package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/steffantucker/AdventofCode2021/helpers"
)

type fish struct {
	age   int
	birth int
}

func main() {
	testinput := helpers.LoadInputLines("testinput")
	inputinput := helpers.LoadInputLines("input")
	fmt.Printf("test 1: %v\n", puzzle1(testinput))
	fmt.Printf("input 1: %v\n", puzzle1(inputinput))
	fmt.Printf("test 2: %v\n", puzzle2(testinput))
	fmt.Printf("input 2: %v\n", puzzle2(inputinput))
}

func positions(in string) (out []int) {
	n := strings.Split(in, ",")
	for _, num := range n {
		i, _ := strconv.Atoi(num)
		out = append(out, i)
	}
	return
}

func minmax(c []int) (min, max int) {
	min, max = c[0], c[1]
	for _, g := range c {
		if min > g {
			min = g
		}
		if max < g {
			max = g
		}
	}
	return
}

func puzzle1(input []string) (out int) {
	crabs := positions(input[0])
	min, max := minmax(crabs)
	out = max * len(crabs)
	for i := min; i <= max; i++ {
		total := 0
		for _, crab := range crabs {
			total += int(math.Abs(float64(crab - i)))
		}
		if total < out {
			out = total
		}
	}
	return
}

func geoSeries(in int) int {
	return (in * (in + 1)) / 2
}

func puzzle2(input []string) (out int) {
	crabs := positions(input[0])
	min, max := minmax(crabs)
	out = max * geoSeries(len(crabs))
	for i := min; i <= max; i++ {
		total := 0
		for _, crab := range crabs {
			dif := int(math.Abs(float64(crab - i)))
			total += geoSeries(dif)
		}
		if total < out {
			out = total
		}
	}
	return
}
