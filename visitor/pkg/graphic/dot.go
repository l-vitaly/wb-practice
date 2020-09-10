package graphic

import "github.com/l-vitaly/wb-practice/visitor/pkg/visitor"

type Dot struct {
	x, y string
}

func (d *Dot) X() string {
	return d.x
}

func (d *Dot) Y() string {
	return d.y
}

func (d *Dot) Accept(v visitor.Visitor) {
	v.VisitDot(d)
}

func NewDot(x string, y string) *Dot {
	return &Dot{x: x, y: y}
}
