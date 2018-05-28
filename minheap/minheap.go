package minheap

type IntHeap []int

func (s IntHeap) Len() int           { return len(s) }
func (s IntHeap) Less(i, j int) bool { return s[i] < s[j] }
func (s IntHeap) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// Push and Pop use pointer receivers
// because they modify the slice's length,
func (s *IntHeap) Push(val interface{}) {
	*s = append(*s, val.(int))
}

// Pop removes the last element and return the new array.
// This is only for sort.Interface.
// Nothing to do with sorting because it just returns the last element.
// Pop in heap does the sorting job.
func (s *IntHeap) Pop() interface{} {
	osl := *s
	n := len(osl)
	ns := osl[n-1]
	*s = osl[0 : n-1]
	return ns
}
