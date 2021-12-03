package sumbarine

type Submarine struct {
	Aim        int
	Horizontal int
	Depth      int
}

func (s Submarine) CalcHash() int {
	return s.Depth * s.Horizontal
}

func (s Submarine) Initialize(d, h, a int) {
	s.Depth = d
	s.Horizontal = h
	s.Aim = a
}
