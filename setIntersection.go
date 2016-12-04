package set

func Intersection(a Set, b Set) (Set, error) {
	// a is countable
	if a.Countable() {
		return finitIntersection(a, b)
	}

	// b is countable
	if b.Countable() {
		return finitIntersection(b, a)
	}

	// Both not countable
	return infiniteIntersection(a, b)
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

func infiniteIntersection(a Set, b Set) (Set, error) {
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
