package set

// Set implements an interface to various implementations of sets.
//
// Countable() implements the indicator for a set to be countable or not.
// Countable set have the ability to count their elements and show them
// explicit in arrays. Therefor countable set implementation should also
// implement Cardinality() and List() since those will be used by functions
// like Difference(a Set, b Set) to be able to speed up the handling of sets.
type Set interface {
	Contains(x interface{}) (bool, error)
	Countable() bool
	Cardinality() (uint64, error)
	List() ([]interface{}, error)
}
