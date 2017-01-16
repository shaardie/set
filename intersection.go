package set

// Intersection creates a set as the intersection of the sets a and b.
//
// If a or b are definitely finite the resulting set is also definitely finite.
// Otherwise the resulting set is not definitely finite.
// So this function is an excellent way to make sets definitely finite.
func Intersection(a Set, b Set) (Set, error) {
	// a is definitely finite
	if a.DefinitelyFinite() {
		return defFiniteIntersection(a, b)
	}
	// b is definitely finite
	if b.DefinitelyFinite() {
		return defFiniteIntersection(b, a)
	}
	// Both not definitely finite
	return notDefFiniteIntersection(a, b)
}

// Creates a new set as an intersection of a and b.
// Here it is assumed that a is definitely finite and therefor a definitely finite set is created.
func defFiniteIntersection(a Set, b Set) (Set, error) {
	// Create new definitely finite set
	newSet := elementSet{make(map[interface{}]struct{})}
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
			newSet.elements[element] = struct{}{}
		}
	}
	return newSet, nil
}

// Creates a new set as an intersection of a and b by using a function as a definer.
// Although this function works on definitely finite sets it is designed to create intersections of two not definitely finite functions.
func notDefFiniteIntersection(a Set, b Set) (Set, error) {
	// both not definitely finite
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
