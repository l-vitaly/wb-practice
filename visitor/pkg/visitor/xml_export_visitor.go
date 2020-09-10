package visitor

type shape interface {
	X() string
	Y() string
}

type dot interface {
	X() string
	Y() string
}

type xmlExportVisitor struct {
	result []string
}

func (v *xmlExportVisitor) VisitShape(shape shape) {
	v.result = append(v.result, "<shape x="+shape.X()+" y="+shape.Y()+" />")
}

func (v *xmlExportVisitor) VisitDot(dot dot) {
	v.result = append(v.result, "<dot x="+dot.X()+" y="+dot.Y()+" />")
}

func (v *xmlExportVisitor) Result() []string {
	return v.result
}

func NewXMLExportVisitor() Visitor {
	return &xmlExportVisitor{}
}
