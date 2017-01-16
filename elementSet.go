package set

type elementSet struct {
	elements map[interface{}]struct{}
}

func (set elementSet) Contains(x interface{}) (bool, error) {
	if _, ok := set.elements[x]; ok {
		return true, nil
	}
	return false, nil
}

func (set elementSet) Countable() bool {
	return set.DefinitelyFinite()
}

func (set elementSet) DefinitelyFinite() bool {
	return true
}

func (set elementSet) Cardinality() (uint64, error) {
	return uint64(len(set.elements)), nil
}

func (set elementSet) List() ([]interface{}, error) {

	number, err := set.Cardinality()
	if err != nil {
		return []interface{}{}, err
	}

	var list = make([]interface{}, 0, number)
	for element, _ := range set.elements {
		list = append(list, element)
	}
	return list, nil
}
