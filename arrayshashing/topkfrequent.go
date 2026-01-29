package arrayshashing

import (
	"container/heap"
	"sort"
)

type datatopk struct {
	num  int
	freq int
}

type pqtopk []*datatopk

var _ heap.Interface = (*pqtopk)(nil)

// Len implements [heap.Interface].
func (pq pqtopk) Len() int {
	return len(pq)
}

// Less implements [heap.Interface].
func (pq pqtopk) Less(i int, j int) bool {
	return pq[i].freq > pq[j].freq // max heap
}

// Pop implements [heap.Interface].
func (pq *pqtopk) Pop() any {
	x := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return x
}

// Push implements [heap.Interface].
func (pq *pqtopk) Push(x any) {
	*pq = append(*pq, x.(*datatopk))
}

// Swap implements [heap.Interface].
func (pq pqtopk) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func TopKFrequent(nums []int, k int) []int {
	res := make([]int, k)
	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
	}

	pq := make(pqtopk, 0, len(m))
	heap.Init(&pq)

	for num, freq := range m {
		data := datatopk{num: num, freq: freq}
		heap.Push(&pq, &data)
	}

	for i := range k {
		res[i] = heap.Pop(&pq).(*datatopk).num
	}

	return res
}

func TopKFrequent2(nums []int, k int) []int {
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	bucket := make([][2]int, 0, len(counter))
	for num, freq := range counter {
		bucket = append(bucket, [2]int{freq, num})
	}

	sort.Slice(bucket, func(i, j int) bool {
		return bucket[i][0] > bucket[j][0]
	})

	res := make([]int, k)
	for i := range k {
		res[i] = bucket[i][1]
	}
	return res
}
