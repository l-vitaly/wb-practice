package graphic

type Dot struct {
	X, Y string
}

func (d *Dot) Accept(visitor visitor) {
	visitor.VisitDot(d)
}
