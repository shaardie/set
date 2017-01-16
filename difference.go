package set

// Difference creates a set as the difference of the sets a and b.
//
// If the set a is definitely finite the resulting set is also definitely finite and if a is not definitely finite the resulting set is not also.
func Difference(a Set, b Set) (Set, error) {
	if a.DefinitelyFinite() {
		return defFiniteDifference(a, b)
	}
	return notDefFiniteDifference(a, b)
}

// The set a is definitely finite.
// Therfore the difference is calculated explicit as a subset of a.
func defFiniteDifference(a Set, b Set) (Set, error) {
	// Create new definitely finite set
	newSet := elementSet{make(map[interface{}]struct{})}

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
			newSet.elements[element] = struct{}{}
		}
	}
	return newSet, nil
}

// The set a is not definitely finite.
// Therefore it is difficult to calculate the difference explicitly.
// We use a function to describe it.
func notDefFiniteDifference(a Set, b Set) (Set, error) {
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
