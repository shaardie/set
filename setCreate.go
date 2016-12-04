package

func createFromArray(list []interface{}) Set {
	set := elementSet{make(map[interface{}]bool)}
	for _, element := range list {
		set.elements[element] = true
	}
	return set
}
