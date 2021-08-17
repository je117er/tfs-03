package main

import (
	"sync"
	"testing"
)

type Person struct {
	Age int
}

var personPool = sync.Pool{
	New: func() interface{} { return new(Person) },
}

func BenchmarkWithoutPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
		p = new(Person)
		p.Age = 23
		}
	}
}

func BenchmarkWithPool(b *testing.B) {
	var p *Person
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 10000; j++ {
		p = personPool.Get().(*Person)
		p.Age = 23
		personPool.Put(p)
		}
	}
}

func BenchmarkPool(b *testing.B) {
	var p sync.Pool
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			p.Put(1)
			p.Get()
		}
	})
}

func BenchmarkAllocation(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i := 0
			i = i
		}
	})
}