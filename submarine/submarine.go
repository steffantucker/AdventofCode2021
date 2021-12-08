package submarine

type Sub struct {
	Aim        int
	Horizontal int
	Depth      int
}

func NewSubmarine() *Sub {
	return &Sub{
		Depth:      0,
		Horizontal: 0,
		Aim:        0,
	}
}

func (s *Sub) CalcHash() int {
	return s.Depth * s.Horizontal
}

func (s *Sub) Initialize(d, h, a int) {
	s.Depth = d
	s.Horizontal = h
	s.Aim = a
}
