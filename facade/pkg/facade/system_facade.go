package facade

type subsystemReady interface {
	OperationReady() string
}

type subsystemGo interface {
	OperationGo() string
}

type subsystemReadyAndGo interface {
	subsystemReady
	subsystemGo
}

type System interface {
	// Operation run big operation.
	Operation() (result string)
}

// System fof operations.
type system struct {
	subsystem1 subsystemReady
	subsystem2 subsystemGo
	subsystem3 subsystemReadyAndGo
}

// Operation run big operation.
func (s *system) Operation() (result string) {
	result = "Facade initializes subsystems:\n"
	result += s.subsystem1.OperationReady()
	result += s.subsystem2.OperationGo()
	result += s.subsystem3.OperationReady()
	result += s.subsystem3.OperationGo()
	return result
}

// NewSystem a creates System instance.
func NewSystem(subsystem1 subsystemReady, subsystem2 subsystemGo, subsystem3 subsystemReadyAndGo) System {
	return &system{subsystem1: subsystem1, subsystem2: subsystem2, subsystem3: subsystem3}
}
