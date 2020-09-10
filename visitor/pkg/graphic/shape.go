package graphic

type Shape interface {
	X() string
	Y() string
}

type shape struct {
	x, y string
}

func (s *shape) X() string {
	return s.x
}

func (s *shape) Y() string {
	return s.y
}

func NewShape(x string, y string) Shape {
	return &shape{x: x, y: y}
}
