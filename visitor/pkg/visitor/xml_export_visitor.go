package visitor

type shape interface {
	X() string
	Y() string
}

type dot interface {
	X() string
	Y() string
}

type XMLExportVisitor struct {
	result []string
}

func (v *XMLExportVisitor) VisitShape(shape shape) {
	v.result = append(v.result, "<shape x="+shape.X()+" y="+shape.Y()+" />")
}

func (v *XMLExportVisitor) VisitDot(dot dot) {
	v.result = append(v.result, "<dot x="+dot.X()+" y="+dot.Y()+" />")
}

func (v *XMLExportVisitor) Result() []string {
	return v.result
}

func NewXMLExportVisitor() *XMLExportVisitor {
	return &XMLExportVisitor{}
}
