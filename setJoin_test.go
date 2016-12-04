package set

import "testing"

func TestCountableJoinBroken(t *testing.T) {
	if _, err := countableJoin(infinitelt3, finite123); err == nil {
		t.Error("No Error")
	}
	if _, err := countableJoin(finite123, infinitegt3); err == nil {
		t.Error("No Error")
	}
	if _, err := countableJoin(infinitegt3, infinitelt3); err == nil {
		t.Error("No Error")
	}
}

func TestCountableJoin(t *testing.T) {
	set, err := countableJoin(finite123, finite345)
	if err != nil {
		t.Fatal(err)
	}
	if !set.Countable() {
		t.Fatal("Set should be countable")
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

func TestNotCountableJoin(t *testing.T) {
	set, err := notCountableJoin(infinitegt3, finite123)
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
	} else if !set.Countable() {
		t.Error("Set not countable")
	}
	if set, err := Join(infinitegt3, infinitegt3); err != nil {
		t.Error(err)
	} else if set.Countable() {
		t.Error("Set countable")
	}
}
