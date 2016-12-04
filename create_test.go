package set

import "testing"

func TestCreateFromArray(t *testing.T) {
	list := []interface{}{1, 2, 3, 4}
	set := CreateFromArray(list)
	if _, ok := set.(elementSet); !ok {
		t.Error("Wrong type")
	}
	for _, element := range list {
		if yes, err := set.Contains(element); err != nil {
			t.Fatal(err)
		} else if !yes {
			t.Errorf("%v not in set", element)
		}
	}
}

func TestCreateFromFunc(t *testing.T) {
	set := CreateFromFunc(func(interface{}) (bool, error) {
		return true, nil
	})
	if _, ok := set.(functionSet); !ok {
		t.Error("Wrong type")
	}
}
