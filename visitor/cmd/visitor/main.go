package main

import (
	"fmt"

	"github.com/l-vitaly/wb-practice/visitor/pkg/graphic"
	"github.com/l-vitaly/wb-practice/visitor/pkg/visitor"
)

func main() {
	allShapes := []graphic.Graphic{
		graphic.NewDot("100", "200"),
		graphic.NewShape("100", "300"),
	}
	v := visitor.NewXMLExportVisitor()
	for _, shape := range allShapes {
		shape.Accept(v)
	}
	for _, s := range v.Result() {
		fmt.Println(s)
	}
}
