package set

import (
	"reflect"
)

var finite123 = elementSet{
	map[interface{}]bool{
		1: true,
		2: true,
		3: true,
	},
}

var finite345 = elementSet{
	map[interface{}]bool{
		3: true,
		4: true,
		5: true,
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
