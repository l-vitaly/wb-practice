package graphic

type dot struct{}

func (c *dot) Value() int {
	return 100
}

func NewDot() *dot {
	return &dot{}
}

func (c *dot) Accept(v interfaceVisitor) {
	v.VisitDot(c)
}
