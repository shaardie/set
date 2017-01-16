package set

import (
	"testing"
)

func TestDefFiniteDifferenceBroken(t *testing.T) {
	_, err := defFiniteDifference(infinitegt3, finite123)
	if err == nil {
		t.Error("First set not countable but no error")
	}
}

func TestDefFiniteDifference(t *testing.T) {
	set, err := defFiniteDifference(finite123, infinitegt3)
	if err != nil {
		t.Fatal(err)
	}
	if !set.DefinitelyFinite() {
		t.Fatal("Resulting set is not definitely finite")
	}
	list, err := set.List()
	if err != nil {
		t.Fatal(err)
	}
	if len(list) == 1 && list[0] == 3 {
		t.Error("Wrong list")
	}
}

func TestNotDefFiniteDifference(t *testing.T) {
	set, err := notDefFiniteDifference(infinitegt3, finite123)
	if err != nil {
		t.Fatal(err)
	}
	if set.DefinitelyFinite() {
		t.Fatal("Resulting set is definitely finite")
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

func TestDifference(t *testing.T) {
	if set, err := Difference(finite123, infinitegt3); err != nil {
		t.Error(err)
	} else if !set.DefinitelyFinite() {
		t.Error("Set not definitely finite")
	}
	if set, err := Difference(infinitegt3, finite123); err != nil {
		t.Error(err)
	} else if set.DefinitelyFinite() {
		t.Error("Set definitely finite")
	}
}
