package mb0_test

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/mb0/critbit"
	"github.com/tv42/benchmark-ordered-map/fixture"
)

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var tree critbit.Tree
		for i := 0; i < len(fixture.TestData); i++ {
			var s string
			p := (*reflect.StringHeader)(unsafe.Pointer(&s))
			p.Data = uintptr(unsafe.Pointer(&fixture.TestData[i]))
			p.Len = int(unsafe.Sizeof(fixture.TestData[0]))
			tree.Insert(s)
		}
	}
}

func BenchmarkSortedInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var tree critbit.Tree
		for i := 0; i < len(fixture.SortedTestData); i++ {
			var s string
			p := (*reflect.StringHeader)(unsafe.Pointer(&s))
			p.Data = uintptr(unsafe.Pointer(&fixture.SortedTestData[i]))
			p.Len = int(unsafe.Sizeof(fixture.SortedTestData[0]))
			tree.Insert(s)
		}
	}
}

func BenchmarkIterate(b *testing.B) {
	var tree critbit.Tree
	for i := 0; i < len(fixture.TestData); i++ {
		var s string
		p := (*reflect.StringHeader)(unsafe.Pointer(&s))
		p.Data = uintptr(unsafe.Pointer(&fixture.TestData[i]))
		p.Len = int(unsafe.Sizeof(fixture.TestData[0]))
		tree.Insert(s)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tree.Iterate("", func(s string) bool {
			item := (*fixture.Item)(unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s)).Data))
			_ = item.Key
			_ = item.Value
			fmt.Printf("got %#v %#v\n", item.Key, item.Value)
			return true
		})
	}
}

// func TestIterate(t *testing.T) {
// 	var tree critbit.Tree
// 	for i := 0; i < len(fixture.TestData); i++ {
// 		var s string
// 		p := (*reflect.StringHeader)(unsafe.Pointer(&s))
// 		p.Data = uintptr(unsafe.Pointer(&fixture.TestData[i]))
// 		p.Len = int(unsafe.Sizeof(fixture.TestData[0]))
// 		tree.Insert(s)
// 	}

// 	tree.Iterate("", func(s string) bool {
// 		item := (*fixture.Item)(unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&s)).Data))
// 		_ = item.Key
// 		_ = item.Value
// 		fmt.Printf("got %#v %#v\n", item.Key, item.Value)
// 		return true
// 	})
// }
