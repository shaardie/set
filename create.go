package set

// CreateFromArray creates Set with `DefinitelyFinite() == true` from an arbitrary list of elements.
func CreateFromArray(list []interface{}) Set {
	set := elementSet{make(map[interface{}]struct{})}
	for _, element := range list {
		set.elements[element] = struct{}{}
	}
	return set
}

// CreateFromFunc creates a Set with `DefinitelyFinite() == false` from a function which indicates if the given element is contained in the set.
//
// To better support complex function it is possible to define an error which is passed through the other functions like Interception to ease error handling.
func CreateFromFunc(f func(interface{}) (bool, error)) Set {
	return functionSet{f}
}
