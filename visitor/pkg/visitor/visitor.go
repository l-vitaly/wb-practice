package visitor

import "fmt"

type xmlExportVisitor struct {
	result []string
}

type dot = interface{ Value() int }
type shape = interface{ Value() int }

func (v *xmlExportVisitor) VisitDot(d dot) {
	v.result = append(v.result, fmt.Sprintf("<shape value=\"%d\" />", d.Value()))
}

func (v *xmlExportVisitor) VisitShape(s shape) {
	v.result = append(v.result, fmt.Sprintf("<shape value=\"%d\" />", s.Value()))
}

func (v *xmlExportVisitor) Result() []string {
	return v.result
}

func NewVisitor() *xmlExportVisitor {
	return &xmlExportVisitor{}
}
