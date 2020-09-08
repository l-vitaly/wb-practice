package subsystem

// Subsystem3
type Subsystem3 struct {
}

func (s *Subsystem3) OperationReady() string {
	return "Subsystem3: Ready!\n"
}

func (s *Subsystem3) OperationGo() string {
	return "Subsystem3: Go!\n"
}
