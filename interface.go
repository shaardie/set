package set

// Set implements an interface to various implementations of sets.
//
// DefinitelyFinite() implements the indicator for a set to be definitely finite.
// But caution: `DefinitlyFinite() == false` does not mean that this set could not be finite, too.
// Definitely finite sets have the ability to count their elements and show them explicit in arrays.
// Therefore, definitely finite set implementations should also implement Cardinality() and List() since those will be used by functions like Difference(a Set, b Set) to be able to speed up the handling of sets.
// Countable() is deprecated. Better use DefinitelyFinite().
type Set interface {
	Contains(x interface{}) (bool, error)
	Countable() bool
	DefinitelyFinite() bool
	Cardinality() (uint64, error)
	List() ([]interface{}, error)
}
