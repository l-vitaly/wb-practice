package builder

// Builder sets.
type Builder interface {
	Reset()
	BuildSet1()
	BuildSet2()
	BuildSet3()
	Result() []string
}
