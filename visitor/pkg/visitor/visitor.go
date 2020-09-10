package visitor

import "github.com/l-vitaly/wb-practice/visitor/pkg/graphic"

type XMLExportVisitor interface {
	VisitShape(shape *graphic.Shape)
	VisitDot(dot *graphic.Dot)
	Result() []string
}

type xmlExportVisitor struct {
	result []string
}

func (v *xmlExportVisitor) VisitShape(shape *graphic.Shape) {
	v.result = append(v.result, "<shape x="+shape.X+" y="+shape.Y+" />")
}

func (v *xmlExportVisitor) VisitDot(dot *graphic.Dot) {
	v.result = append(v.result, "<dot x="+dot.X+" y="+dot.Y+" />")
}

func (v *xmlExportVisitor) Result() []string {
	return v.result
}

func NewXMLExportVisitor() XMLExportVisitor {
	return &xmlExportVisitor{}
}
