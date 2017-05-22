package main

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
		// lookup[string(find)] = i
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
