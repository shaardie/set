package set

// createFromArray creates a countable Set from an arbitrary list of elements
func createFromArray(list []interface{}) Set {
	set := elementSet{make(map[interface{}]bool)}
	for _, element := range list {
		set.elements[element] = true
	}
	return set
}

// createFromFunc creates a non countable Set from an function which indicates
// if the given element is contained in the set
func createFromFunc(f func(interface{}) (bool, error)) Set {
	return functionSet{f}
}
