package set

func Difference(a Set, b Set) (Set, error) {

	// a is countable
	if a.Countable() {
		newSet := elementSet{make(map[interface{}]bool)}

		elements, err := a.List()
		if err != nil {
			return newSet, err
		}

		for element := range elements {
			if inB, err := b.Contains(element); err != nil {
				return newSet, err
			} else if !inB { // not in b, so in difference
				newSet.elements[element] = true
			}
		}
		return newSet, nil
	}

	// If a is not countable
	newContains := func(x interface{}) (bool, error) {
		if inA, err := a.Contains(x); err != nil {
			return false, err
		} else if !inA {
			return false, nil
		}

		// If here, x is in a
		if inB, err := b.Contains(x); err != nil {
			return false, err
		} else if inB {
			return false, nil
		}

		// if here, x not in b
		return true, nil

	}
	return functionSet{newContains}, nil

}

func Intersection(a Set, b Set) (Set, error) {
	// a is countable
	if a.Countable() {
		return finitIntersection(a, b)
	}

	// b is countable
	if b.Countable() {
		return finitIntersection(b, a)
	}

	// both not countable
	newContains := func(x interface{}) (bool, error) {
		if inA, err := a.Contains(x); err != nil {
			return false, err
		} else if !inA {
			return false, nil
		}

		// If here, x is in a
		if inB, err := b.Contains(x); err != nil {
			return false, err
		} else if !inB {
			return false, nil
		}

		// if here, x is in b
		return true, nil

	}
	return functionSet{newContains}, nil

}

func finitIntersection(a Set, b Set) (Set, error) {
	newSet := elementSet{}

	elements, err := a.List()
	if err != nil {
		return newSet, err
	}

	for _, element := range elements {
		inB, err := b.Contains(element)
		if err != nil {
			return newSet, err
		}
		if inB {
			newSet.elements[element] = true
		}
	}
	return newSet, nil
}

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
