package cuckoo

import (
	"fmt"
	"testing"
)

func test_filter(t* testing.T)  {

	c := NewCuckoo(10, .1)
	c.insert(newId("hello"))
	c.insert(newId("world"))
	ok := c.lookup(newId("world"))
	fmt.Printf("%v\n", ok)
	c.delete(newId("world"))
	ok = c.lookup(newId("world"))
	fmt.Printf("%v\n", ok)
}
