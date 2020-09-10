package graphic

type interfaceVisitor interface {
	VisitDot(d interface{ Value() int })
	VisitShape(s interface{ Value() int })
}

type Graphic interface {
	Accept(visitor interfaceVisitor)
}
