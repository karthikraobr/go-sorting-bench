package sorting

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"

	"github.com/shawnsmithdev/zermelo"
)

func foo() []int {
	hold := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		a := rand.Intn(100000000)
		hold = append(hold, a)
	}
	return hold
}

func DefaultSort(a []int) {
	sort.Ints(a)
}

func ZormeloSort(a []int) {
	zermelo.Sort(a)
}

func filterAndSortDefault(a []int) {
	visited := make(map[int]struct{}, 0)
	accumulator := make([]int, 0)
	for i := 0; i < 10; i++ {
		for _, val := range a {
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
	for i := 0; i < 10; i++ {
		for _, val := range a {
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
	accumulator := make([]int, 0)
	for i := 0; i < 10; i++ {
		for _, val := range a {
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
	for i := 0; i < 10; i++ {
		for _, val := range a {
			if _, ok := visited[val]; !ok {
				accumulator = append(accumulator, val)
				visited[val] = struct{}{}
			}
		}
		zermelo.Sort(accumulator)
	}

}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
