package sortedSlice_test

import (
	"sort"
	"testing"

	"github.com/tv42/benchmark-ordered-map/fixture"
)

type SortedSlice []fixture.Item

func (n *SortedSlice) Add(d fixture.Item) {
	i := sort.Search(len(*n), func(i int) bool {
		return (*n)[i].Key >= d.Key
	})
	if i < len(*n) && (*n)[i].Key == d.Key {
		// replace
		(*n)[i] = d
	} else {
		// not yet present, insert at i

		// let append worry about reallocation; dummy placeholder
		r := append(*n, fixture.Item{})

		// shuffle to insert
		copy((r)[i+1:], (r)[i:])
		(r)[i] = d
		*n = r
	}
}

func BenchmarkInsert(b *testing.B) {
	var list *SortedSlice
	for i := 0; i < b.N; i++ {
		list = &SortedSlice{}
		for i := 0; i < len(fixture.TestData); i++ {
			list.Add(fixture.TestData[i])
		}
	}

	// verify
	b.StopTimer()
	started := false
	var seen fixture.Key
	for i := 0; i < len(*list); i++ {
		k := (*list)[i].Key
		if started && seen >= k {
			b.Errorf("Key wrong order: got %v but seen %v", k, seen)
		}
		seen = k
		started = true
		got := (*list)[i]
		want := fixture.FinalTestData[k]
		if got != want {
			b.Errorf("Key wrong value: got %v want %v", got, want)
		}
	}
}

func BenchmarkSortedInsert(b *testing.B) {
	var list *SortedSlice
	for i := 0; i < b.N; i++ {
		list = &SortedSlice{}
		for i := 0; i < len(fixture.SortedTestData); i++ {
			list.Add(fixture.SortedTestData[i])
		}
	}

	// verify
	b.StopTimer()
	started := false
	var seen fixture.Key
	for i := 0; i < len(*list); i++ {
		k := (*list)[i].Key
		if started && seen >= k {
			b.Errorf("Key wrong order: got %v but seen %v", k, seen)
		}
		seen = k
		started = true
		got := (*list)[i]
		want := fixture.FinalTestData[k]
		if got != want {
			b.Errorf("Key wrong value: got %v want %v", got, want)
		}
	}
}

func BenchmarkIterate(b *testing.B) {
	list := &SortedSlice{}
	for i := 0; i < len(fixture.TestData); i++ {
		list.Add(fixture.TestData[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := 0; i < len(*list); i++ {
			_ = (*list)[i].Key
			_ = (*list)[i].Value
		}
	}

	// verify
	b.StopTimer()
	started := false
	var seen fixture.Key
	for i := 0; i < len(*list); i++ {
		k := (*list)[i].Key
		if started && seen >= k {
			b.Errorf("Key wrong order: got %v but seen %v", k, seen)
		}
		seen = k
		started = true
		got := (*list)[i]
		want := fixture.FinalTestData[k]
		if got != want {
			b.Errorf("Key wrong value: got %v want %v", got, want)
		}
	}
}
