package graphic

type Dot interface {
	X() string
	Y() string
}

type dot struct {
	x, y string
}

func (d *dot) X() string {
	return d.x
}

func (d *dot) Y() string {
	return d.y
}

func NewDot(x string, y string) Dot {
	return &dot{x: x, y: y}
}
