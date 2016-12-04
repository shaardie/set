package set

// CreateFromArray creates a countable Set from an arbitrary list of elements.
//
// Example:
//	CreateFromArray([]interface{}{ true, "string", 42 })
func CreateFromArray(list []interface{}) Set {
	set := elementSet{make(map[interface{}]bool)}
	for _, element := range list {
		set.elements[element] = true
	}
	return set
}

// CreateFromFunc creates a non countable Set from a function which indicates
// if the given element is contained in the set.
//
// To better support complex function it is possible to define an error which
// is passed through the other functions like Interception to ease error
// handling.
//
// Example:
//	CreateFromFunc(func(interface{}) (bool, error)) Set { false, nil }) // empty set
func CreateFromFunc(f func(interface{}) (bool, error)) Set {
	return functionSet{f}
}
