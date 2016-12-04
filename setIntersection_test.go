package set

import (
	"testing"
)

func TestCountableIntersectionBroken(t *testing.T) {
	if _, err := countableIntersection(infinitegt3, infinitegt3); err == nil {
		t.Error("First set not countable but no error")
	}
}

func TestCountableIntersection(t *testing.T) {
	set, err := countableIntersection(finite123, infinitegt3)
	if err != nil {
		t.Fatal(err)
	}

	if !set.Countable() {
		t.Fatal("Set should be countable")
	}

	if list, err := set.List(); err != nil {
		t.Fatal(err)
	} else if !(len(list) == 1 && list[0] == 3) {
		t.Error("Wrong set")
	}
}

func TestNotCountableIntersection(t *testing.T) {
	set, err := notCountableIntersection(infinitegt3, infinitelt3)
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
	} else if !set.Countable() {
		t.Error("Set not countable")
	}
	if set, err := Intersection(infinitegt3, finite123); err != nil {
		t.Error(err)
	} else if !set.Countable() {
		t.Error("Set not countable")
	}
	if set, err := Intersection(infinitegt3, infinitegt3); err != nil {
		t.Error(err)
	} else if set.Countable() {
		t.Error("Set countable")
	}
}
