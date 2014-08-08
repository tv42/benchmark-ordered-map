package fixture

import (
	"math"
	"math/rand"
	"sort"
)

type Key int64

// github.com/dataence/skiplist does not seem to know how to iterate
// all values without explicit min and max. Try not to need these.
const (
	KeyMin Key = math.MinInt64
	KeyMax Key = math.MaxInt64
)

type Value struct {
	Gibberish int64
	Filler    int64
}

// Item combines a Key and a Value. Container libraries which store
// the key separately should use Key and Value, and libraries which
// store the whole item in one should use Item.
//
// Benchmarks must NOT use pointers, as that just makes Value be
// pointer-sized, and we can simulate that just fine with changing the
// type Value here.
type Item struct {
	Key
	Value
}

const Count = 100000

var TestData [Count]Item
var SortedTestData [Count]Item
var FinalTestData map[Key]Item

type itemsSort []Item

func (a itemsSort) Len() int {
	return len(a)
}
func (a itemsSort) Less(i, j int) bool {
	return a[i].Key < a[j].Key
}
func (a itemsSort) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func init() {
	FinalTestData = make(map[Key]Item)
	r := rand.New(rand.NewSource(42))
	for i := 0; i < len(TestData); i++ {
		TestData[i].Key = Key(r.Int63())
		TestData[i].Gibberish = r.Int63()
		TestData[i].Filler = r.Int63()

		FinalTestData[TestData[i].Key] = TestData[i]
	}

	copy(SortedTestData[:], TestData[:])
	sort.Sort(itemsSort(SortedTestData[:]))
}
