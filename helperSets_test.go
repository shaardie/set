package set

import (
	"reflect"
)

var finite123 = elementSet{
	map[interface{}]struct{}{
		1: struct{}{},
		2: struct{}{},
		3: struct{}{},
	},
}

var finite345 = elementSet{
	map[interface{}]struct{}{
		3: struct{}{},
		4: struct{}{},
		5: struct{}{},
	},
}

var infinitegt3 = functionSet{
	func(x interface{}) (bool, error) {
		if v := reflect.ValueOf(x); v.Kind() == reflect.Int && v.Int() >= 3 {
			return true, nil
		}
		return false, nil
	},
}

var infinitelt3 = functionSet{
	func(x interface{}) (bool, error) {
		if v := reflect.ValueOf(x); v.Kind() == reflect.Int && v.Int() <= 3 {
			return true, nil
		}
		return false, nil
	},
}
