package main

import (
	"fmt"

	"github.com/steffantucker/AdventofCode2021/helpers"
	"github.com/steffantucker/AdventofCode2021/submarine"
)

func main() {
	testinput := helpers.LoadInputLines("testinput")
	inputinput := helpers.LoadInputLines("input")

	p1test := puzzle1(testinput)
	p1 := puzzle1(inputinput)
	p2test := puzzle2(testinput)
	p2 := puzzle2(inputinput)
	fmt.Printf("test 1: %v\ninput 1: %v\ntest 2: %v\ninput 2: %v\n", p1test, p1, p2test, p2)
}

func puzzle1(input []string) int {
	prog := submarine.NewSubmarineProgram(input)
	prog.RegisterCommand("forward", Forward).RegisterCommand("up", Up).RegisterCommand("down", Down)
	return prog.Run()
}

func Forward(s *submarine.Sub, dist int) {
	s.Horizontal += dist
}

func Up(s *submarine.Sub, dist int) {
	s.Depth -= dist
}

func Down(s *submarine.Sub, dist int) {
	s.Depth += dist
}

func puzzle2(input []string) int {
	prog := submarine.NewSubmarineProgram(input)
	prog.RegisterCommand("forward", Forward2).RegisterCommand("up", Up2).RegisterCommand("down", Down2)
	return prog.Run()
}

func Forward2(s *submarine.Sub, dist int) {
	s.Horizontal += dist
	s.Depth += s.Aim * dist
}

func Up2(s *submarine.Sub, dist int) {
	s.Aim -= dist
}

func Down2(s *submarine.Sub, dist int) {
	s.Aim += dist
}
