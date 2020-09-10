package main

import (
	"fmt"

	"github.com/l-vitaly/wb-practice/visitor/pkg/graphic"
	"github.com/l-vitaly/wb-practice/visitor/pkg/visitor"
)

func main() {
	allGraphics := []graphic.Graphic{
		graphic.NewShape(),
		graphic.NewDot(),
	}
	v := visitor.NewVisitor()
	for _, g := range allGraphics {
		g.Accept(v)
	}
	for _, s := range v.Result() {
		fmt.Println(s)
	}
}
