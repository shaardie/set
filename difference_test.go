package set

import (
	"testing"
)

func TestCountableDifferenceBroken(t * testing.T) {
	_, err := countableDifference(infinitegt3, finite123)
	if err == nil {
		t.Error("First set not countable but no error")
	}
}

func TestCountableDifference(t *testing.T) {
	set, err := countableDifference(finite123, infinitegt3)
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

func TestNotCountableDifference(t *testing.T) {
	set, err := notCountableDifference(infinitegt3, finite123)
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

func TestDifference(t*testing.T) {
	if set, err := Difference(finite123, infinitegt3); err != nil {
		t.Error(err)
	} else if !set.Countable() {
		t.Error("Set not countable")
	}
	if set, err := Difference(infinitegt3, finite123); err != nil {
		t.Error(err)
	} else if set.Countable() {
		t.Error("Set countable")
	}
}
