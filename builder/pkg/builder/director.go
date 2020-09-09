package builder

// Director for management builder.
type Director interface {
	// Make sets.
	Make()
}

type director struct {
	builder Builder
}

// Make sets.
func (d *director) Make() {
	d.builder.Reset()
	d.builder.BuildSet1()
	d.builder.BuildSet2()
	d.builder.BuildSet3()
}

// NewDirector a creates Director instance.
func NewDirector(builder Builder) Director {
	return &director{builder: builder}
}
