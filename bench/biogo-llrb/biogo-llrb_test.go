package biogollrb_test

import (
	"testing"

	"github.com/biogo/store/llrb"
	"github.com/tv42/benchmark-ordered-map/fixture"
)

type llrbItemUnique fixture.Item

func (i llrbItemUnique) Compare(b llrb.Comparable) int {
	return int(i.Key - b.(llrbItemUnique).Key)
}

type llrbItemNonUnique fixture.Item

func (i llrbItemNonUnique) Compare(b llrb.Comparable) int {
	d := int(i.Key - b.(llrbItemNonUnique).Key)
	if d == 0 {
		return 1 // Insert after.
	}
	return d
}

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var tree llrb.Tree
		for i := 0; i < len(fixture.TestData); i++ {
			tree.Insert(llrbItemUnique(fixture.TestData[i]))
		}
	}
}

func BenchmarkSortedInsert_ReplaceOrInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var tree llrb.Tree
		for i := 0; i < len(fixture.SortedTestData); i++ {
			tree.Insert(llrbItemUnique(fixture.SortedTestData[i]))
		}
	}
}

func BenchmarkSortedInsert_InsertNoReplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var tree llrb.Tree
		for i := 0; i < len(fixture.SortedTestData); i++ {
			tree.Insert(llrbItemNonUnique(fixture.SortedTestData[i]))
		}
	}
}

func BenchmarkIterate(b *testing.B) {
	var tree llrb.Tree
	for i := 0; i < len(fixture.TestData); i++ {
		tree.Insert(llrbItemUnique(fixture.TestData[i]))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tree.Do(func(i llrb.Comparable) bool {
			_ = i.(llrbItemUnique).Key
			_ = i.(llrbItemUnique).Value
			return false
		})
	}
}
