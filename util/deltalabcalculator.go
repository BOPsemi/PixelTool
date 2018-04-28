package util

/*
DeltaLabCalculator :interface of delta lab calculator
*/
type DeltaLabCalculator interface {
}

// deltaLabCalculater str definition
type deltaLabCalculator struct {
}

/*
NewDeltaLabCalculator : initializer
*/
func NewDeltaLabCalculator() DeltaLabCalculator {
	obj := new(deltaLabCalculator)

	return obj
}
