package ryszardSkiplist_test

import (
	"testing"

	"github.com/ryszard/goskiplist/skiplist"
	"github.com/tv42/benchmark-ordered-map/fixture"
)

func lt(l, r interface{}) bool {
	return l.(fixture.Key) < r.(fixture.Key)
}

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := skiplist.NewCustomMap(lt)
		for i := 0; i < len(fixture.TestData); i++ {
			list.Set(fixture.TestData[i].Key, fixture.TestData[i].Value)
		}
	}
}

func BenchmarkSortedInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := skiplist.NewCustomMap(lt)
		for i := 0; i < len(fixture.SortedTestData); i++ {
			list.Set(fixture.SortedTestData[i].Key, fixture.SortedTestData[i].Value)
		}
	}
}

func BenchmarkIterate(b *testing.B) {
	list := skiplist.NewCustomMap(lt)
	for i := 0; i < len(fixture.TestData); i++ {
		list.Set(fixture.TestData[i].Key, fixture.TestData[i].Value)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for i := list.Iterator(); i.Next(); {
			_ = i.Key().(fixture.Key)
			_ = i.Value().(fixture.Value)
		}
	}
}
