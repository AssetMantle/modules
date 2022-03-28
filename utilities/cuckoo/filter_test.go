package cuckoo

import (
	"fmt"
	"testing"
)

func test_filter(t *testing.T) {

	c := NewCuckoo(10, .1)
	c.Insert(NewID("hello"))
	c.Insert(NewID("world"))
	ok := c.Lookup(NewID("world"))
	fmt.Printf("%v\n", ok)
	c.Delete(NewID("world"))
	ok = c.Lookup(NewID("world"))
	fmt.Printf("%v\n", ok)
}
