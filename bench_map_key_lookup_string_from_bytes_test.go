package main

/*
Results
--
$ go test -v -bench BenchmarkMap -benchmem
testing: warning: no tests to run
BenchmarkMapLookupKeyStringFromBytes-4          100000000               19.6 ns/op             0 B/op          0 allocs/op
BenchmarkMapSetKeyStringFromBytes-4             20000000                73.8 ns/op             5 B/op          1 allocs/op
PASS
ok      github.com/robskillington/benchmarks-go 3.561s
*/

import (
	"fmt"
	"testing"
)

func BenchmarkMapLookupKeyStringFromBytes(b *testing.B) {
	entries := 4096
	lookup := make(map[string]int, entries)
	for i := 0; i < entries; i++ {
		lookup[fmt.Sprintf("foo.%d", i)] = -1
	}

	find := []byte("foo.0")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, ok := lookup[string(find)]; !ok {
			b.Fatalf("key %s should exist", string(find))
		}
	}
}

func BenchmarkMapSetKeyStringFromBytes(b *testing.B) {
	entries := 4096
	lookup := make(map[string]int, entries)
	for i := 0; i < entries; i++ {
		lookup[fmt.Sprintf("foo.%d", i)] = -1
	}

	find := []byte("foo.0")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		lookup[string(find)] = i
	}
}
