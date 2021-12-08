package submarine

import (
	"strconv"
	"strings"
)

type (
	SubmarineProgram struct {
		sub      *Sub
		instr    []Instructions
		commands map[string]CommandFunction
	}

	CommandFunction func(*Sub, int)
)

func NewSubmarineProgram(in []string) *SubmarineProgram {
	s := &SubmarineProgram{
		sub:      NewSubmarine(),
		commands: make(map[string]CommandFunction, 3),
	}
	s.LoadInstructions(in)
	return s
}

func (s *SubmarineProgram) RegisterCommand(command string, f CommandFunction) *SubmarineProgram {
	s.commands[command] = f
	return s
}

func (s *SubmarineProgram) InitSubmarine(d, h, a int) {
	s.sub.Initialize(d, h, a)
}

func (s *SubmarineProgram) LoadInstructions(in []string) {
	for _, i := range in {
		parts := strings.Split(i, " ")
		c := parts[0]
		n, _ := strconv.Atoi(parts[1])
		s.instr = append(s.instr, Instructions{Distance: n, Command: c})
	}
}

func (s *SubmarineProgram) Run() int {
	for _, i := range s.instr {
		if com, ok := s.commands[i.Command]; ok {
			com(s.sub, i.Distance)
		}
	}
	return s.sub.CalcHash()
}
