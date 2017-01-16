package set

import "testing"

func TestDefFiniteJoinBroken(t *testing.T) {
	if _, err := defFiniteJoin(infinitelt3, finite123); err == nil {
		t.Error("No Error")
	}
	if _, err := defFiniteJoin(finite123, infinitegt3); err == nil {
		t.Error("No Error")
	}
	if _, err := defFiniteJoin(infinitegt3, infinitelt3); err == nil {
		t.Error("No Error")
	}
}

func TestDefFiniteJoin(t *testing.T) {
	set, err := defFiniteJoin(finite123, finite345)
	if err != nil {
		t.Fatal(err)
	}
	if !set.DefinitelyFinite() {
		t.Fatal("Set should be definitely finite")
	}

	if number, err := set.Cardinality(); err != nil {
		t.Fatal(err)
	} else if number != 5 {
		t.Error("Wrong number of elements")
	}

	for i := 1; i < 6; i++ {
		if yes, err := set.Contains(i); err != nil {
			t.Fatal(err)
		} else if !yes {
			t.Errorf("%v not in set", i)
		}
	}
}

func TestNotDefFiniteJoin(t *testing.T) {
	set, err := notDefFiniteJoin(infinitegt3, finite123)
	if err != nil {
		t.Fatal(err)
	}
	for i := 1; i < 999; i++ {
		if yes, err := set.Contains(i); err != nil {
			t.Fatal(err)
		} else if !yes {
			t.Errorf("%v not in set", i)
		}
	}
}

func TestJoin(t *testing.T) {
	if set, err := Join(finite123, finite123); err != nil {
		t.Error(err)
	} else if !set.DefinitelyFinite() {
		t.Error("Set not definitely finite")
	}
	if set, err := Join(infinitegt3, infinitegt3); err != nil {
		t.Error(err)
	} else if set.DefinitelyFinite() {
		t.Error("Set definitely finite")
	}
}
