package set

type elementSet struct {
	elements map[interface{}]bool
}

func (set elementSet) Contains(x interface{}) (bool, error) {
	return set.elements[x], nil
}

func (set elementSet) Countable() bool {
	return true
}

func (set elementSet) Cardinality() (uint64, error) {

	var number uint64
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
