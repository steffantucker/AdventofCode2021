package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/steffantucker/AdventofCode2021/helpers"
)

var charPairs = map[string]string{"{": "}", "(": ")", "[": "]", "<": ">"}

const (
	opening = "([{<"
	closing = ")]}>"
)

func main() {
	testinput := helpers.LoadInputLines("testinput")
	inputinput := helpers.LoadInputLines("input")
	fmt.Printf("test 1: %v\n", puzzle1(testinput))
	fmt.Printf("input 1: %v\n", puzzle1(inputinput))
	fmt.Printf("test 2: %v\n", puzzle2(testinput))
	fmt.Printf("input 2: %v\n", puzzle2(inputinput))
}

func puzzle1(input []string) (out int) {
	for _, row := range input {
		in := strings.Split(row, "")
		p := traverse(in, []string{})
		out += p
	}
	return
}

func traverse(rest []string, unclosed []string) (point int) {
	if len(rest) == 0 {
		return 0
	}
	char := rest[0]
	if strings.Contains(opening, char) {
		unclosed = append(unclosed, char)
		return traverse(rest[1:], unclosed)
	}
	if strings.Contains(closing, char) {
		if char != charPairs[unclosed[len(unclosed)-1]] {
			return points(char)
		} else {
			unclosed = unclosed[:len(unclosed)-1]
		}
	}
	return traverse(rest[1:], unclosed)
}

func puzzle2(input []string) (out int) {
	allpoints := []int{}
	for _, row := range input {
		in := strings.Split(row, "")
		p, fix := traverse2(in, []string{})
		if p == 0 {
			allpoints = append(allpoints, points2(fix))
		}
	}
	sort.Ints(allpoints)
	middle := math.Floor(float64(len(allpoints)) / 2.0)
	return allpoints[int(middle)]
}

func traverse2(rest []string, unclosed []string) (point int, fix []string) {
	if len(rest) == 0 {
		return 0, unclosed
	}
	char := rest[0]
	if strings.Contains(opening, char) {
		unclosed = append(unclosed, char)
		return traverse2(rest[1:], unclosed)
	}
	if strings.Contains(closing, char) {
		if char != charPairs[unclosed[len(unclosed)-1]] {
			return points(char), []string{}
		} else {
			unclosed = unclosed[:len(unclosed)-1]
		}
	}
	return traverse2(rest[1:], unclosed)
}

func points2(fix []string) (out int) {
	for i := len(fix) - 1; i >= 0; i-- {
		char := fix[i]
		p := 0
		switch charPairs[char] {
		case ")":
			p = 1
		case "]":
			p = 2
		case "}":
			p = 3
		case ">":
			p = 4
		}
		out = out*5 + p
	}
	return
}

func points(expected string) int {
	switch expected {
	case ")":
		return 3
	case "]":
		return 57
	case "}":
		return 1197
	case ">":
		return 25137
	default:
		panic(fmt.Sprintf("bad char: %v\n", expected))
	}
}
