package set

type Set interface {
	Contains(x interface{}) (bool, error)
	Countable() (bool, error)
	Cardinality() (int64, error)
	List() ([]interface{}, error)
	Difference(s Set) (Set, error)
	Intersection(s Set) (Set, error)
	Join(s Set) (Set, error)
}
