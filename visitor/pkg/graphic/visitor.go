package graphic

type Graphic interface {
	Accept(visitor visitor)
}

type visitor interface {
	VisitShape(shape *Shape)
	VisitDot(dot *Dot)
	Result() []string
}
