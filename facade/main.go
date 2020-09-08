package main

import (
	"fmt"

	"github.com/l-vitaly/wb-practice/facade/pkg/facade"
	"github.com/l-vitaly/wb-practice/facade/pkg/subsystem"
)

func main() {
	subsystem1 := new(subsystem.Subsystem1)
	subsystem2 := new(subsystem.Subsystem2)
	subsystem3 := new(subsystem.Subsystem3)

	system := facade.NewSystem(subsystem1, subsystem2, subsystem3)

	fmt.Print(system.Operation())
}
