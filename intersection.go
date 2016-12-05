package set

// Intersection creates a set as the intersection of the sets a and b.
//
// If a and b are countable the resulting set is also countable otherwise the
// resulting set is not countable. So this function is an excellent way to make
// not countable sets countable.
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

// Create a new set as an intersection of a and b. Here it is assumed that a is
// countable and therefor a countable set is created.
func countableIntersection(a Set, b Set) (Set, error) {
	// Create new countable set
	newSet := elementSet{make(map[interface{}]bool)}
	// Explicit list of the elements of a
	elements, err := a.List()
	if err != nil {
		return newSet, err
	}
	// Check if the elements are also in b and therefor in the intersection
	for _, element := range elements {
		yes, err := b.Contains(element)
		if err != nil {
			return newSet, err
		}
		if yes {
			newSet.elements[element] = true
		}
	}
	return newSet, nil
}

// Create a new set as an intersection of a and b by using a function as a
// definer. Although this function work on countable sets it is designed to
// create intersection of two not countable functions.
func notCountableIntersection(a Set, b Set) (Set, error) {
	// both not countable
	newContains := func(x interface{}) (bool, error) {
		if yes, err := a.Contains(x); err != nil {
			return false, err
		} else if !yes {
			return false, nil
		}
		// If here, x is in a
		if yes, err := b.Contains(x); err != nil {
			return false, err
		} else if !yes {
			return false, nil
		}
		// if here, x is in b
		return true, nil
	}
	return functionSet{newContains}, nil
}
