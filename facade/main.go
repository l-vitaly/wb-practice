package main

import (
	"fmt"

	"github.com/l-vitaly/wb-practice/facade/pkg/facade"
	"github.com/l-vitaly/wb-practice/facade/pkg/subsystem"
)

func main() {
	subsystem1 := subsystem.NewSubsystem1()
	subsystem2 := subsystem.NewSubsystem2()
	subsystem3 := subsystem.NewSubsystem3()

	system := facade.NewSystem(subsystem1, subsystem2, subsystem3)

	fmt.Print(system.Operation())
}
