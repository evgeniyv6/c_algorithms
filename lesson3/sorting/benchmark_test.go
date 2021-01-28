package sorting

import (
	"github.com/evgeniyv6/c_algorithms/lesson3/slices"
	"testing"
	)


var (
	GlobalRes    int
	sliceElemNum = 10
)

func BenchmarkShakerSorting(b *testing.B) {
	var res int
	mySl := slices.GenerSlice(sliceElemNum)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res = ShakerSorting(&mySl)
	}
	GlobalRes = res
}

func BenchmarkShakerSortingFlag(b *testing.B) {
	var res int
	mySl := slices.GenerSlice(sliceElemNum)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res = ShakerSortingFlag(&mySl)
	}
	GlobalRes = res
}

func BenchmarkShakerSortingMemPos(b *testing.B) {
	var res int
	mySl := slices.GenerSlice(sliceElemNum)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		res = ShakerSortingMemPos(&mySl)
	}
	GlobalRes = res
}


// Parallel template
//func BenchmarkShakerSortingMemPos(b *testing.B) {
//	mySl := slices.GenerSlice(sliceElemNum)
//	b.ResetTimer()
//	b.RunParallel(func(pb *testing.PB) {
//		var buf bytes.Buffer
//		for pb.Next() {
//			buf.Reset()
//			ShakerSortingMemPos(&mySl)
//		}
//	})
//}
