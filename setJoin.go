package set

func Join(a Set, b Set) (Set, error) {

	// if a and b countable
	if a.Countable() && b.Countable() {
		newSet := elementSet{}

		for _, set := range []Set{a, b} {
			list, err := set.List()
			if err != nil {
				return newSet, err
			}
			for _, element := range list {
				newSet.elements[element] = true
			}
		}
		return newSet, nil
	}

	newContains := func(x interface{}) (bool, error) {
		for _, set := range []Set{a, b} {
			yes, err := set.Contains(x)
			if err != nil {
				return false, err
			}
			if yes {
				return true, nil
			}
		}
		return false, nil
	}
	return functionSet{newContains}, nil
}
