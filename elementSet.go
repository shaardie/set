package set

import (
	"errors"
)

type elementSet struct {
	elements map[interface{}]bool
}

func (set elementSet) Contains(x interface{}) (bool, error) {
	return set.elements[x], nil
}

func (set elementSet) Countable() (bool, error) {
	return true, nil
}

func (set elementSet) Cardinality() (int64, error) {

	var number int64
	for _, contained := range set.elements {
		if contained {
			number++
		}
	}

	return number, nil
}

func (set elementSet) List() ([]interface{}, error) {

	number, err := set.Cardinality()
	if err != nil {
		return []interface{}{}, nil
	}

	var list = make([]interface{}, 0, number)
	for element, contained := range set.elements {
		if contained {
			list = append(list, element)
		}
	}
	return list, nil
}

func (set elementSet) Difference(s Set) (Set, error) {
	newSet := elementSet{}

	for element, inSet := range set.elements {
		if inS, err := s.Contains(element); err != nil {
			return newSet, err
		} else if inS {
			continue
		}
		if inSet {
			newSet.elements[element] = true
		}
	}
	return newSet, nil
}

func (set elementSet) Intersection(s Set) (Set, error) {
	newSet := elementSet{}

	for element, inSet := range set.elements {
		if inS, err := s.Contains(element); err != nil {
			return newSet, err
		} else if inS && inSet {
			newSet.elements[element] = true
		}
	}
	return newSet, nil
}

func (set elementSet) Join(s Set) (Set, error) {
	if countable, err := s.Countable(); err != nil {
		return elementSet{}, err
	} else if countable {
		newSet := elementSet{}
		list, err := s.List()
		if err != nil {
			return newSet, err
		}
		for element := range list {
			newSet.elements[element] = true
		}
		for element, contained := range set.elements {
			if contained {
				newSet.elements[element] = true
			}
		}
		return newSet, nil
	}
	return elementSet{}, errors.New("Unable to Join")
}
