package glennbrownSkiplist_test

import (
	"testing"

	"github.com/glenn-brown/skiplist"
	"github.com/tv42/benchmark-ordered-map/fixture"
)

type glennBrownKey fixture.Key

func (a glennBrownKey) Less(b interface{}) bool {
	return a < b.(glennBrownKey)
}

func (k glennBrownKey) Score() float64 {
	return float64(k)
}

func BenchmarkInsertIterate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := skiplist.New()
		for i := 0; i < len(fixture.TestData); i++ {
			list.Set(glennBrownKey(fixture.TestData[i].Key), fixture.TestData[i].Value)
		}

		e := list.Front()
		for e != nil {
			_ = fixture.Key(e.Key().(glennBrownKey))
			_ = e.Value.(fixture.Value)
			e = e.Next()
		}
	}
}

func BenchmarkSortedInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := skiplist.New()
		for i := 0; i < len(fixture.SortedTestData); i++ {
			list.Set(glennBrownKey(fixture.SortedTestData[i].Key), fixture.SortedTestData[i].Value)
		}
	}
}
