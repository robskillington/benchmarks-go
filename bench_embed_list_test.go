package main

import "testing"

type list struct {
	root listElem
}

type listElem struct {
	next *listElem
	prev *listElem
}

type entity struct {
	listElem
	x int
}

func BenchmarkListOfEntity(b *testing.B) {
	l := &list{}
	l.root.next = &l.root
	l.root.prev = &l.root

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		initList(l)
	}
}

func initList(l *list) {
	x, y, z := &entity{x: 1}, &entity{x: 2}, &entity{x: 3}

	l.root.next = &x.listElem
	x.next = &y.listElem
	y.next = &z.listElem
}
