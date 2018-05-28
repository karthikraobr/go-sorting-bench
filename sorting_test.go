package sorting

import (
	"testing"
)

func BenchmarkEverything(b *testing.B) {
	tt := []struct {
		name string
		f    func([]int)
	}{
		{name: "sortDefault", f: DefaultSort},
		{name: "sortZermelo", f: ZormeloSort},
		{name: "filterAndSortOnceDefault", f: filterAndSortDefault},
		{name: "filterAndSortOnceZermelo", f: filterAndSortZermelo},
		{name: "FilterAndContinuousSortDefault", f: filterAndContinuousSortDefault},
		{name: "FilterAndContinuousSortZermelo", f: filterAndContinuousSortZermelo},
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
