package set

import (
	"testing"
)

func TestDifferenceFirstCountable(t *testing.T) {
	set, err := Difference(finite123, infinitegt3)
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
	set, err := Difference(infinitegt3, finite123)
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
