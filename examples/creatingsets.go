package main

import (
	"github.com/shaardie/set"
)

func main() {

	// Creating a set explicit from a list of elements
	set.CreateFromArray([]interface{}{true, "string", 42})

  // Creating a set from a function
  set.CreateFromFunc(func(interface{}) (Set, error) { return false, nil }) // empty set
}
