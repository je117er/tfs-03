package main

import "sync"

type Person struct {
	Name string
}

// initializes pool
var personPool = sync.Pool{
	// returns a new object if pool is empty
	// when called by personPool.Put()
	New: func() interface{} { return new(Person) },
}

func main() {
	// gets a new instance
	newPerson := personPool.Get().(*Person)

	// defers release function so the instance
	// can be used again
	defer personPool.Put(newPerson)

	// uses the instance
	newPerson.Name = "Jack"
}


