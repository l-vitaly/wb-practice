package main

import (
	"fmt"

	"github.com/l-vitaly/wb-practice/builder/pkg/builder"
	"github.com/l-vitaly/wb-practice/builder/pkg/concretebuilder"
)

func main() {
	concreteBuilder := concretebuilder.NewConcreteBuilder()
	director := builder.NewDirector(concreteBuilder)
	director.Make()

	for _, s := range concreteBuilder.Result() {
		fmt.Println(s)
	}
}
