package set

// Join creates a set as the join of the sets a and b.
//
// If a and b are countable the resulting set is also countable otherwise the
// resulting set not countable.
func Join(a Set, b Set) (Set, error) {
	if a.Countable() && b.Countable() {
		return countableJoin(a, b)
	}
	return notCountableJoin(a, b)
}

// Creates a new set as an join of a and b. Here is assumed that a and b are
// countable.
func countableJoin(a Set, b Set) (Set, error) {
  // Create a countable set
	newSet := elementSet{make(map[interface{}]bool)}
  // Loop to add all elements from a and b explicit
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

// Creates a new set as the join of a and b by using a function to define the
// set. Therefor the resulting set is not countable.
func notCountableJoin(a Set, b Set) (Set, error) {
	newContains := func(x interface{}) (bool, error) {
    // Loop to return true for all elements contains in a or b
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
