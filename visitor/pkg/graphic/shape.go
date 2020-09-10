package graphic

import "github.com/l-vitaly/wb-practice/visitor/pkg/visitor"

type Shape struct {
	x, y string
}

func (s *Shape) X() string {
	return s.x
}

func (s *Shape) Y() string {
	return s.y
}

func (s *Shape) Accept(v visitor.Visitor) {
	v.VisitShape(s)
}

func NewShape(x string, y string) *Shape {
	return &Shape{x: x, y: y}
}
