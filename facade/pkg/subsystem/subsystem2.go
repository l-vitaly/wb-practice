package subsystem

// Subsystem2
type Subsystem2 struct {
}

func (s *Subsystem2) OperationGo() string {
	return "Subsystem2: Go!\n"
}

func NewSubsystem2() *Subsystem2 {
	return &Subsystem2{}
}
