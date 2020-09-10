package main

import (
	"fmt"

	"github.com/l-vitaly/wb-practice/visitor/pkg/graphic"

	"github.com/l-vitaly/wb-practice/visitor/pkg/visitor"
)

func main() {
	allShapes := []graphic.Graphic{
		&graphic.Dot{
			X: "100",
			Y: "200",
		},
		&graphic.Shape{
			X: "300",
			Y: "100",
		},
	}
	v := visitor.NewXMLExportVisitor()
	for _, shape := range allShapes {
		shape.Accept(v)
	}
	for _, s := range v.Result() {
		fmt.Println(s)
	}
}
