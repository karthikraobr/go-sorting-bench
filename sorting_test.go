package sorting

import (
	"testing"
)

func BenchmarkEverything(b *testing.B) {
	tt := []struct {
		name string
		f    func([]int)
	}{
		{name: "filterAndSortOnceDefault", f: filterAndSortDefault},
		{name: "filterAndSortOnceZermelo", f: filterAndSortZermelo},
		{name: "filterAndContinuousSortDefault", f: filterAndContinuousSortDefault},
		{name: "filterAndContinuousSortZermelo", f: filterAndContinuousSortZermelo},
		{name: "heap", f: filterAndSortHeap},
		{name: "btree", f: filterAndSortTree},
	}

	a := foo()
	for _, tc := range tt {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tc.f(a)
			}
		})
	}
}
