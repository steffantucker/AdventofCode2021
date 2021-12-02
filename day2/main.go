package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/steffantucker/AdventofCode2021/helpers"
)

type submarine struct {
	depth int
	aim   int
	horiz int
}

type instruction struct {
	dir  int
	dist int
}

func main() {
	testinput := helpers.LoadInputLines("testinput")
	inputinput := helpers.LoadInputLines("input")

	p1test := puzzle1(testinput)
	p1 := puzzle1(inputinput)
	p2test := puzzle2(testinput)
	p2 := puzzle2(inputinput)
	fmt.Printf("test 1: %v\ninput 1: %v\ntest 2: %v\ninput 2: %v\n", p1test.calcHash(), p1.calcHash(), p2test.calcHash(), p2.calcHash())
}

func (s submarine) calcHash() int {
	return s.depth * s.horiz
}

func puzzle1(input []string) (out submarine) {
	for _, v := range input {
		parts := strings.Split(v, " ")
		dir := parts[0]
		dist, _ := strconv.Atoi(parts[1])
		switch dir {
		case "forward":
			out.horiz += dist
		case "up":
			out.depth -= dist
		case "down":
			out.depth += dist
		}
	}
	return
}

func puzzle2(input []string) (out submarine) {
	for _, v := range input {
		parts := strings.Split(v, " ")
		dir := parts[0]
		dist, _ := strconv.Atoi(parts[1])
		switch dir {
		case "forward":
			out.horiz += dist
			out.depth += (out.aim * dist)
		case "up":
			out.aim -= dist
		case "down":
			out.aim += dist
		}
	}
	return
}
