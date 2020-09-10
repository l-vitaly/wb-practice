package visitor

import "fmt"

type xmlExportVisitor struct {
	result []string
}

func (v *xmlExportVisitor) VisitDot(d interface{ Value() int }) {
	v.result = append(v.result, fmt.Sprintf("<shape value=\"%d\" />", d.Value()))
}

func (v *xmlExportVisitor) VisitShape(s interface{ Value() int }) {
	v.result = append(v.result, fmt.Sprintf("<shape value=\"%d\" />", s.Value()))
}

func (v *xmlExportVisitor) Result() []string {
	return v.result
}

func NewVisitor() *xmlExportVisitor {
	return &xmlExportVisitor{}
}
