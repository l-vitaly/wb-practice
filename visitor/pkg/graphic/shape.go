package graphic

type shape struct{}

func (c *shape) Value() int {
	return 400
}

func NewShape() *shape {
	return &shape{}
}

func (c *shape) Accept(v interfaceVisitor) {
	v.VisitShape(c)
}
