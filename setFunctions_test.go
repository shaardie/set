package set

import (
	"reflect"
	"testing"
)

var a = elementSet{
	map[interface{}]bool{
		1: true,
		2: true,
		3: true,
	},
}

var b = elementSet{
	map[interface{}]bool{
		3: true,
		4: true,
		5: true,
	},
}

var c = functionSet{
	func(x interface{}) (bool, error) {
		if v := reflect.ValueOf(x); v.Kind() == reflect.Int && v.Int() <= 3 {
			return true, nil
		}
		return false, nil
	},
}

var d = functionSet{
	func(x interface{}) (bool, error) {
		if v := reflect.ValueOf(x); v.Kind() == reflect.Int && v.Int() >= 3 {
			return true, nil
		}
		return false, nil
	},
}

func TestDifferenceFirstCountable(t *testing.T) {
	set, err := Difference(a, d)
	if err != nil {
		t.Fatal(err)
	}
	if !set.Countable() {
		t.Fatal("Resulting set is not countable")
	}
	list, err := set.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(list) == 1 && list[0] == 3 {
		t.Error("Wrong list")
	}
}

func TestDifferenceFirstNotCountable(t *testing.T) {
	set, err := Difference(d, a)
	if err != nil {
		t.Fatal(err)
	}
	if set.Countable() {
		t.Fatal("Resulting set is countable")
	}
	for _, element := range []int{1, 2, 3} {
		yes, err := set.Contains(element)
		if err != nil {
			t.Error(err)
			continue
		}
		if yes {
			t.Errorf("%v in set", element)
		}
	}
	for _, element := range []int{4, 5, 6} {
		yes, err := set.Contains(element)
		if err != nil {
			t.Error(err)
			continue
		}
		if !yes {
			t.Errorf("%v not in set", element)
		}
	}
}
