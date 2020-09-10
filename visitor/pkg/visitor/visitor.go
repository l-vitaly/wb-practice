package visitor

type Visitor interface {
	VisitShape(shape shape)
	VisitDot(dot dot)
	Result() []string
}
