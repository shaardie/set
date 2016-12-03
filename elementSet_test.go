package set

import (
	"testing"
)

func TestElementSetContains(t *testing.T) {
	elements := map[interface{}]bool{
		"string": true,
		true:     false,
		false:    true,
		42:       true,
	}

	set := elementSet{elements}

	for _, element := range []interface{}{"string", 42, false} {
		if contained, err := set.Contains(element); err != nil {
			t.Fatal(err)
		} else if !contained {
			t.Errorf("%v is not containd", element)
		}
	}

	for _, element := range []interface{}{true, "No in Set"} {
		if contained, err := set.Contains(element); err != nil {
			t.Fatal(err)
		} else if contained {
			t.Errorf("%v is containd", element)
		}
	}
}

func TestElementSetCountable(t *testing.T) {
	set := elementSet{}
	if !set.Countable() {
		t.Error("elementSet not countable...ridiculous!")
	}
}

func TestElementSetCardinality(t *testing.T) {
	check_number := func(set elementSet, t *testing.T, should uint64) {
		if number, err := set.Cardinality(); err != nil {
			t.Fatal(err)
		} else if number != should {
			t.Errorf("Empty set has not cardinality %v", should)
		}

	}

	set := elementSet{}
	check_number(set, t, 0)

	set.elements = map[interface{}]bool{1: true, 2: true, 3: true}
	check_number(set, t, 3)
}
