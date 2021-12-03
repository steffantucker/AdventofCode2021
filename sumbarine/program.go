package sumbarine

import (
	"strconv"
	"strings"
)

type (
	SubmarineProgram struct {
		sub      Submarine
		instr    []Instructions
		commands map[string]CommandFunction
	}

	CommandFunction func(*Submarine, int)
)

func (s SubmarineProgram) RegisterCommand(command string, f CommandFunction) {
	s.commands[command] = f
}

func (s SubmarineProgram) InitSubmarine(d, h, a int) {
	s.sub.Initialize(d, h, a)
}

func (s SubmarineProgram) LoadInstructions(in []string) {
	for _, i := range in {
		parts := strings.Split(i, " ")
		c := parts[0]
		n, _ := strconv.Atoi(parts[1])
		s.instr = append(s.instr, Instructions{Distance: n, Command: c})
	}
}
