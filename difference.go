package set

// Create a set as the difference of the sets a and b.
//
// If the set a is countable the resulting set is also countable and if a is
// not countable the resulting set is not also.
func Difference(a Set, b Set) (Set, error) {
	if a.Countable() {
		return countableDifference(a, b)
	}
	return notCountableDifference(a, b)
}

// The set a is countable. Therfor the difference is calculated explicit as a
// subset of a.
func countableDifference(a Set, b Set) (Set, error) {
	// Create new countable set
	newSet := elementSet{make(map[interface{}]bool)}

	// Explicit list of elements in a
	elements, err := a.List()
	if err != nil {
		return newSet, err
	}

	// exclude all elements also in b
	for element := range elements {
		if yes, err := b.Contains(element); err != nil {
			return newSet, err
		} else if !yes {
			newSet.elements[element] = true
		}
	}
	return newSet, nil
}

// The set a is not countable. Therefor it is difficult to calculat the
// difference explicit. We use a function to describe it.
func notCountableDifference(a Set, b Set) (Set, error) {
	newContains := func(x interface{}) (bool, error) {
		if yes, err := a.Contains(x); err != nil {
			return false, err
		} else if !yes {
			return false, nil
		}

		// If here, x is in a
		if yes, err := b.Contains(x); err != nil {
			return false, err
		} else if yes {
			return false, nil
		}

		// if here, x not in b
		return true, nil

	}
	return functionSet{newContains}, nil
}
