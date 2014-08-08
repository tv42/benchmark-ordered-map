package petarGoLLRB_test

import (
	"testing"

	"github.com/petar/GoLLRB/llrb"
	"github.com/tv42/benchmark-ordered-map/fixture"
)

type llrbItem fixture.Item

func (a llrbItem) Less(b llrb.Item) bool {
	return a.Key < b.(llrbItem).Key
}

func BenchmarkInsertIterate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := llrb.New()
		for i := 0; i < len(fixture.TestData); i++ {
			tree.ReplaceOrInsert(llrbItem(fixture.TestData[i]))
		}

		min := tree.Min()
		tree.AscendGreaterOrEqual(min, func(i llrb.Item) bool {
			_ = i.(llrbItem).Key
			_ = i.(llrbItem).Value
			return true
		})
	}
}

func BenchmarkSortedInsert_ReplaceOrInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := llrb.New()
		for i := 0; i < len(fixture.SortedTestData); i++ {
			tree.ReplaceOrInsert(llrbItem(fixture.SortedTestData[i]))
		}
	}
}

func BenchmarkSortedInsert_InsertNoReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tree := llrb.New()
		for i := 0; i < len(fixture.SortedTestData); i++ {
			tree.InsertNoReplace(llrbItem(fixture.SortedTestData[i]))
		}
	}
}
