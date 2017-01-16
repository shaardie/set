package set

import (
	"testing"
)

func TestDefFiniteIntersectionBroken(t *testing.T) {
	if _, err := defFiniteIntersection(infinitegt3, infinitegt3); err == nil {
		t.Error("First set not definitely finite but no error")
	}
}

func TestDefFiniteIntersection(t *testing.T) {
	set, err := defFiniteIntersection(finite123, infinitegt3)
	if err != nil {
		t.Fatal(err)
	}

	if !set.DefinitelyFinite() {
		t.Fatal("Set should be definitely finite")
	}

	if list, err := set.List(); err != nil {
		t.Fatal(err)
	} else if !(len(list) == 1 && list[0] == 3) {
		t.Error("Wrong set")
	}
}

func TestNotDefFiniteIntersection(t *testing.T) {
	set, err := notDefFiniteIntersection(infinitegt3, infinitelt3)
	if err != nil {
		t.Fatal(err)
	}

	if yes, err := set.Contains(3); err != nil {
		t.Fatal(err)
	} else if !yes {
		t.Error("3 not in set")
	}

	for _, element := range []int{1, 2, 4, 5, 6} {
		if yes, err := set.Contains(element); err != nil {
			t.Fatal(err)
		} else if yes {
			t.Errorf("%v in set", element)
		}
	}
}

func TestIntersection(t *testing.T) {
	if set, err := Intersection(finite123, infinitegt3); err != nil {
		t.Error(err)
	} else if !set.DefinitelyFinite() {
		t.Error("Set not definitely finite")
	}
	if set, err := Intersection(infinitegt3, finite123); err != nil {
		t.Error(err)
	} else if !set.DefinitelyFinite() {
		t.Error("Set not definitely finite")
	}
	if set, err := Intersection(infinitegt3, infinitegt3); err != nil {
		t.Error(err)
	} else if set.DefinitelyFinite() {
		t.Error("Set definitely finite")
	}
}
