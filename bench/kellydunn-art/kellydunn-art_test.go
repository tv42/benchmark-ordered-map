package kellydunnART_test

import (
	"encoding/binary"
	"testing"

	"github.com/kellydunn/go-art"
	"github.com/tv42/benchmark-ordered-map/fixture"
)

func makeKey(k fixture.Key, buf []byte) {
	u := uint64(k) ^ 1<<63
	binary.BigEndian.PutUint64(buf, u)
}

func BenchmarkInsert(b *testing.B) {
	key := make([]byte, 8)
	for i := 0; i < b.N; i++ {
		tree := art.NewArtTree()
		for i := 0; i < len(fixture.TestData); i++ {
			makeKey(fixture.TestData[i].Key, key)
			tree.Insert(key, fixture.TestData[i])
		}
	}
}

func BenchmarkSortedInsert(b *testing.B) {
	key := make([]byte, 8)
	for i := 0; i < b.N; i++ {
		tree := art.NewArtTree()
		for i := 0; i < len(fixture.SortedTestData); i++ {
			makeKey(fixture.SortedTestData[i].Key, key)
			tree.Insert(key, fixture.SortedTestData[i])
		}
	}
}

func BenchmarkIterate(b *testing.B) {
	key := make([]byte, 8)
	tree := art.NewArtTree()
	for i := 0; i < len(fixture.TestData); i++ {
		makeKey(fixture.TestData[i].Key, key)
		tree.Insert(key, fixture.TestData[i])
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// only provides callback-based iteration; meh api
		tree.Each(func(n *art.ArtNode) {
			if !n.IsLeaf() {
				return
			}
			item := n.Value().(fixture.Item)
			_ = item.Key
			_ = item.Value
		})
	}
}
