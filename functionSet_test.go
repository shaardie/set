package set

import (
	"errors"
	"testing"
)

func containsAll(interface{}) (bool, error) {
	return true, nil
}

func containsNone(interface{}) (bool, error) {
	return false, nil
}

func containsErrors(interface{}) (bool, error) {
	return false, errors.New("This is always wrong")
}

func TestFunctionSetContains(t *testing.T) {

	elements := []interface{}{1, 2, 3, 4, "Hello", "World", "!", true, false}

	a := functionSet{containsAll}
	b := functionSet{containsNone}
	c := functionSet{containsErrors}

	for _, element := range elements {

		if yes, err := a.Contains(element); err != nil {
			t.Fatal(err)
		} else if !yes {
			t.Error("%v should be in a", element)
		}

		if yes, err := b.Contains(element); err != nil {
			t.Fatal(err)
		} else if yes {
			t.Error("%v should not be in a", element)
		}

		if _, err := c.Contains(element); err == nil {
			t.Error("Empty error")
		}
	}
}

func TestFunctionSetCountable(t *testing.T) {
	a := functionSet{}
	if a.Countable() {
		t.Error("This should not be countable")
	}
}

func TestFunctionSetCardinality(t *testing.T) {
	a := functionSet{}
	if _, err := a.Cardinality(); err == nil {
		t.Error("Error should not be nil")
	}
}

func TestFunctionSetList(t *testing.T) {
	a := functionSet{}
	if _, err := a.List(); err == nil {
		t.Error("Error should not be nil")
	}
}
