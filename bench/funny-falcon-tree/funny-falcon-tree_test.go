package funnyFalconTree_test

import (
	"testing"

	"github.com/funny-falcon/go-tree"
	"github.com/tv42/benchmark-ordered-map/fixture"
)

type SortedSlice []fixture.Item

func (s SortedSlice) Len() int {
	return len(s)
}

func (s SortedSlice) Less(i, j int) bool {
	return s[i].Key < s[j].Key
}

func (s SortedSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type TreeSlice struct {
	S SortedSlice
	I tree.Tree
}

func (t *TreeSlice) Add(d fixture.Item) {
	i := t.I.Search(func(i int) bool {
		return t.S[i].Key >= d.Key
	})
	if i < len(t.S) && t.S[i].Key == d.Key {
		t.S[i] = d
	} else {
		t.S = append(t.S, d)
		t.I.Insert(t.S)
	}
}

func (t *TreeSlice) AddSorted(d fixture.Item) {
	if len(t.S) > 0 && t.S[len(t.S)-1].Key == d.Key {
		t.S[len(t.S)-1] = d
	} else {
		t.S = append(t.S, d)
	}
}

func (t *TreeSlice) FinishAddSorted() {
	t.I.InitSorted(len(t.S))
}

func (t *TreeSlice) Next(i int) int {
	return t.I.Next(i)
}

func (t *TreeSlice) Clear() {
	t.S = nil
	t.I = tree.Tree{}
}

func BenchmarkInsert(b *testing.B) {
	var list TreeSlice
	for i := 0; i < b.N; i++ {
		list.Clear()
		for i := 0; i < len(fixture.TestData); i++ {
			list.Add(fixture.TestData[i])
		}
	}

	// verify
	b.StopTimer()
	started := false
	var seen fixture.Key
	for i := list.Next(-1); i < len(list.S); i = list.Next(i) {
		k := list.S[i].Key
		if started && seen >= k {
			b.Errorf("Key wrong order: got %v but seen %v", k, seen)
		}
		seen = k
		started = true
		got := list.S[i]
		want := fixture.FinalTestData[k]
		if got != want {
			b.Errorf("Key wrong value: got %v want %v", got, want)
		}
	}
}

func BenchmarkSortedInsert_Regular(b *testing.B) {
	var list TreeSlice
	for i := 0; i < b.N; i++ {
		list.Clear()
		for i := 0; i < len(fixture.SortedTestData); i++ {
			list.Add(fixture.SortedTestData[i])
		}
	}

	// verify
	b.StopTimer()
	started := false
	var seen fixture.Key
	for i := list.Next(-1); i < len(list.S); i = list.Next(i) {
		k := list.S[i].Key
		if started && seen >= k {
			b.Errorf("Key wrong order: got %v but seen %v", k, seen)
		}
		seen = k
		started = true
		got := list.S[i]
		want := fixture.FinalTestData[k]
		if got != want {
			b.Errorf("Key wrong value: got %v want %v", got, want)
		}
	}
}

func BenchmarkSortedInsert_Special(b *testing.B) {
	var list TreeSlice
	for i := 0; i < b.N; i++ {
		list.Clear()
		for i := 0; i < len(fixture.SortedTestData); i++ {
			list.AddSorted(fixture.SortedTestData[i])
		}
		list.FinishAddSorted()
	}

	// verify
	b.StopTimer()
	started := false
	var seen fixture.Key
	for i := list.Next(-1); i < len(list.S); i = list.Next(i) {
		k := list.S[i].Key
		if started && seen >= k {
			b.Errorf("Key wrong order: got %v but seen %v", k, seen)
		}
		seen = k
		started = true
		got := list.S[i]
		want := fixture.FinalTestData[k]
		if got != want {
			b.Errorf("Key wrong value: got %v want %v", got, want)
		}
	}
}

func BenchmarkIterate(b *testing.B) {
	list := TreeSlice{}
	for i := 0; i < len(fixture.TestData); i++ {
		list.Add(fixture.TestData[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := list.Next(-1); i < len(list.S); i = list.Next(i) {
			_ = list.S[i].Key
			_ = list.S[i].Value
		}
	}

	// verify
	b.StopTimer()
	started := false
	var seen fixture.Key
	for i := list.Next(-1); i < len(list.S); i = list.Next(i) {
		k := list.S[i].Key
		if started && seen >= k {
			b.Errorf("Key wrong order: got %v but seen %v", k, seen)
		}
		seen = k
		started = true
		got := list.S[i]
		want := fixture.FinalTestData[k]
		if got != want {
			b.Errorf("Key wrong value: got %v want %v", got, want)
		}
	}
}
