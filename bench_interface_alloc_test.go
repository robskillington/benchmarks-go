package main

import "testing"

var existingFooBar = &fooBar{}

type Foo interface {
	Foo()
}

type FooBar interface {
	Foo()
	Bar()
}

type fooBar struct {
	x int
}

func (*fooBar) Foo() {}
func (*fooBar) Bar() {}

type thing struct {
	foo Foo
}

func newThing(foo Foo) thing {
	return thing{foo: foo}
}

func BenchmarkIfaceAllocs(b *testing.B) {
	ref := FooBar(existingFooBar)
	for i := 0; i < b.N; i++ {
		x := newThing(ref)
		if x.foo != ref {
			b.Fail()
		}
		ref = &fooBar{x: i}
	}
}
