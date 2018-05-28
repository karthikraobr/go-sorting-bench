package sorting

import (
	"container/heap"
	"math/rand"
	"sort"

	btree "github.com/google/btree"
	"github.com/karthikraobr/sorting/minheap"
	"github.com/shawnsmithdev/zermelo"
)

func foo() []int {
	hold := make([]int, 0, 1000000)
	for i := 0; i < 1000000; i++ {
		hold = append(hold, random(1, 1000000))
	}
	return hold
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func filterAndSortDefault(a []int) {
	visited := make(map[int]struct{}, 0)
	accumulator := make([]int, 0)
	var x, y int
	for i := 0; i < 10; i++ {
		x = 1 + y
		y = 100000 * (i + 1)
		for _, val := range a[x:y] {
			if _, ok := visited[val]; !ok {
				accumulator = append(accumulator, val)
				visited[val] = struct{}{}
			}
		}

	}
	sort.Ints(accumulator)
}

func filterAndSortZermelo(a []int) {
	visited := make(map[int]struct{}, 0)
	accumulator := make([]int, 0)

	var x, y int
	for i := 0; i < 10; i++ {
		x = 1 + y
		y = 100000 * (i + 1)
		for _, val := range a[x:y] {
			if _, ok := visited[val]; !ok {
				accumulator = append(accumulator, val)
				visited[val] = struct{}{}
			}
		}

	}
	zermelo.Sort(accumulator)
}

func filterAndContinuousSortDefault(a []int) {
	visited := make(map[int]struct{}, 0)
	var x, y int
	accumulator := make([]int, 0)
	for i := 0; i < 10; i++ {
		x = 1 + y
		y = 100000 * (i + 1)
		for _, val := range a[x:y] {
			if _, ok := visited[val]; !ok {
				accumulator = append(accumulator, val)
				visited[val] = struct{}{}
			}
		}
		sort.Ints(accumulator)
	}
}

func filterAndContinuousSortZermelo(a []int) {
	visited := make(map[int]struct{}, 0)
	accumulator := make([]int, 0)
	var x, y int
	for i := 0; i < 10; i++ {
		x = 1 + y
		y = 100000 * (i + 1)
		for _, val := range a[x:y] {
			if _, ok := visited[val]; !ok {
				accumulator = append(accumulator, val)
				visited[val] = struct{}{}
			}
		}
		zermelo.Sort(accumulator)
	}
}

func filterAndContinuousSortHeap(a []int) {
	visited := make(map[int]struct{}, 0)
	h := &minheap.IntHeap{}
	heap.Init(h)
	var x, y int
	for i := 0; i < 10; i++ {
		x = 1 + y
		y = 100000 * (i + 1)
		for _, val := range a[x:y] {
			if _, ok := visited[val]; !ok {
				heap.Push(h, val)
				visited[val] = struct{}{}
			}
		}
	}
	res := make([]int, 0)
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(int))
	}
}

func filterAndContinuousSortTree(a []int) {
	tr := btree.New(100)
	var x, y int
	for i := 0; i < 10; i++ {
		x = 1 + y
		y = 100000 * (i + 1)
		for _, val := range a[x:y] {
			tr.ReplaceOrInsert(btree.Int(val))
		}
	}
	accumulator := make([]int, 0, tr.Len())
	tr.Ascend(func(a btree.Item) bool {
		bla := a.(btree.Int)
		accumulator = append(accumulator, int(bla))
		return true
	})
}
