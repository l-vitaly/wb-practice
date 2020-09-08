package subsystem

// Subsystem1
type Subsystem1 struct {
}

func (s *Subsystem1) OperationReady() string {
	return "Subsystem1: Ready!\n"
}

func NewSubsystem1() *Subsystem1 {
	return &Subsystem1{}

}
