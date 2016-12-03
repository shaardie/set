package set

import (
	"errors"
)

type functionSet struct {
	contains func(interface{}) (bool, error)
}

func (set functionSet) Contains(x interface{}) (bool, error) {
	return set.contains(x)
}

func (set functionSet) Countable() bool {
	return false
}

func (set functionSet) Cardinality() (uint64, error) {
	return 0, errors.New("Not countable by design")
}

func (set functionSet) List() ([]interface{}, error) {
	return []interface{}{}, errors.New("Not listable by design")
}
