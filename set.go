package set

type Set interface {
	Contains(x interface{}) (bool, error)
	Countable() bool
	Cardinality() (uint64, error)
	List() ([]interface{}, error)
}
