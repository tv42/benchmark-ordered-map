package huanduSkiplist_test

import (
	"testing"

	"github.com/huandu/skiplist"
	"github.com/tv42/benchmark-ordered-map/fixture"
)

type huandySpec struct{}

func (huandySpec) Descending() bool {
	return false
}

func (huandySpec) Compare(a, b interface{}) bool {
	return a.(fixture.Key) < b.(fixture.Key)
}

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := skiplist.New(huandySpec{})
		for i := 0; i < len(fixture.TestData); i++ {
			list.Set(fixture.TestData[i].Key, fixture.TestData[i].Value)
		}
	}
}

func BenchmarkSortedInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := skiplist.New(huandySpec{})
		for i := 0; i < len(fixture.SortedTestData); i++ {
			list.Set(fixture.SortedTestData[i].Key, fixture.SortedTestData[i].Value)
		}
	}
}

func BenchmarkIterate(b *testing.B) {
	list := skiplist.New(huandySpec{})
	for i := 0; i < len(fixture.TestData); i++ {
		list.Set(fixture.TestData[i].Key, fixture.TestData[i].Value)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		e := list.Front()
		for e != nil {
			_ = e.Key().(fixture.Key)
			_ = e.Value.(fixture.Value)
			e = e.Next()
		}
	}
}
