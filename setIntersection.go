package set

func Intersection(a Set, b Set) (Set, error) {
	// a is countable
	if a.Countable() {
		return countableIntersection(a, b)
	}

	// b is countable
	if b.Countable() {
		return countableIntersection(b, a)
	}

	// Both not countable
	return notCountableIntersection(a, b)
}

func countableIntersection(a Set, b Set) (Set, error) {
	newSet := elementSet{make(map[interface{}]bool)}

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

func notCountableIntersection(a Set, b Set) (Set, error) {
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
