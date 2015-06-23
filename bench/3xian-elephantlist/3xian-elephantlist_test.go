package threexianElephantlist_test

import (
	"testing"

	"github.com/3xian/elephantlist"
	"github.com/tv42/benchmark-ordered-map/fixture"
)

func threexianLT(a, b interface{}) bool {
	ka := a.(fixture.Key)
	kb := b.(fixture.Key)
	return ka < kb
}

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := elephantlist.NewList(threexianLT)
		for i := 0; i < len(fixture.TestData); i++ {
			tree.Set(fixture.TestData[i].Key, fixture.TestData[i].Value)
		}
	}
}

func BenchmarkSortedInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := elephantlist.NewList(threexianLT)
		for i := 0; i < len(fixture.SortedTestData); i++ {
			tree.Set(fixture.SortedTestData[i].Key, fixture.SortedTestData[i].Value)
		}
	}
}

// TODO has no iteration
