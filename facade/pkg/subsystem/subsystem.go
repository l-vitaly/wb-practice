package subsystem

// Structures of a third party framework. we
// we have no control over this code, so we cannot simplify it.

// Subsystem1
type Subsystem1 struct {
}

func (s *Subsystem1) OperationReady() string {
	return "Subsystem1: Ready!\n"
}

// Subsystem2
type Subsystem2 struct {
}

func (s *Subsystem2) OperationGo() string {
	return "Subsystem2: Go!\n"
}

// Subsystem3
type Subsystem3 struct {
}

func (s *Subsystem3) OperationReady() string {
	return "Subsystem3: Ready!\n"
}

func (s *Subsystem3) OperationGo() string {
	return "Subsystem3: Go!\n"
}
