package graphic

type Shape struct {
	X, Y string
}

func (s *Shape) Accept(visitor visitor) {
	visitor.VisitShape(s)
}
