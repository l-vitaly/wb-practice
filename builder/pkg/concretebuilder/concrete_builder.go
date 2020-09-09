package concretebuilder

// ConcreteBuilder strings print.
type ConcreteBuilder struct {
	results []string
}

// Reset strings slice.
func (b *ConcreteBuilder) Reset() {
	b.results = []string{}
}

// BuildSet1 run set1.
func (b *ConcreteBuilder) BuildSet1() {
	b.results = append(b.results, "Build Set 1")
}

// BuildSet1 run set2.
func (b *ConcreteBuilder) BuildSet2() {
	b.results = append(b.results, "Build Set 2")
}

// BuildSet1 run set3.
func (b *ConcreteBuilder) BuildSet3() {
	b.results = append(b.results, "Build Set 3")
}

// Result builder.
func (b *ConcreteBuilder) Result() []string {
	return b.results
}

// NewConcreteBuilder a creates ConcreteBuilder instance.
func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{}
}
