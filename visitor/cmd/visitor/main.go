package main

import (
	"fmt"

	"github.com/l-vitaly/wb-practice/visitor/pkg/graphic"
	"github.com/l-vitaly/wb-practice/visitor/pkg/visitor"
)

func main() {
	v := visitor.NewXMLExportVisitor()
	v.VisitDot(graphic.NewDot("100", "200"))
	v.VisitShape(graphic.NewShape("200", "400"))

	for _, s := range v.Result() {
		fmt.Println(s)
	}
}
