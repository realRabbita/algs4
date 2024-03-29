package sort

import (
	"errors"
	"fmt"

	"algs4/typ"
)

// IndexMinPQ Indexed Minimum Priority Queue.
type IndexMinPQ struct {
	maxN int
	n    int
	keys []typ.Comparable // 输入的原始数组
	pq   []int            // sorted index in priority queue (sorted index -> original index)
	qp   []int            // original index -> sorted index
}

func NewIndexMinPQWithCap(initCap int) *IndexMinPQ {
	pq, qp := make([]int, initCap+1), make([]int, initCap+1)
	for i := 0; i < initCap+1; i++ {
		qp[i] = -1 // -1 indicates no such element
	}
	return &IndexMinPQ{
		maxN: initCap,
		n:    0,
		keys: make([]typ.Comparable, initCap+1),
		pq:   pq,
		qp:   qp,
	}
}

func NewIndexMinPQWithKeys(keys []typ.Comparable) *IndexMinPQ {
	minPQ := NewIndexMinPQWithCap(len(keys))
	for i, key := range keys {
		minPQ.Insert(i, key)
	}
	return minPQ
}

func (minPQ IndexMinPQ) validateIndex(i int) error {
	if i < 0 || i >= minPQ.maxN {
		return errors.New("out of bound")
	}
	return nil
}

func (minPQ IndexMinPQ) Size() int {
	return minPQ.n
}

func (minPQ IndexMinPQ) IsEmpty() bool {
	return minPQ.n == 0
}

func (minPQ IndexMinPQ) Contains(i int) bool {
	err := minPQ.validateIndex(i)
	if err != nil {
		fmt.Println("Contains error:", err)
		return false
	}
	return minPQ.qp[i] != -1
}

func (minPQ *IndexMinPQ) Insert(i int, key typ.Comparable) {
	err := minPQ.validateIndex(i)
	if err != nil {
		fmt.Println("Insert error:", err)
		return
	}
	if minPQ.Contains(i) {
		fmt.Println("Insert error: index is already in the priority queue")
		return
	}
	minPQ.n++
	minPQ.qp[i] = minPQ.n
	minPQ.pq[minPQ.n] = i
	minPQ.keys[i] = key
	minPQ.swim(minPQ.n)
}

func (minPQ *IndexMinPQ) Change(i int, key typ.Comparable) {
	err := minPQ.validateIndex(i)
	if err != nil {
		fmt.Println("Change error:", err)
		return
	}
	if !minPQ.Contains(i) {
		fmt.Println("Change error: index is not existed in the priority queue")
		return
	}
	minPQ.keys[i] = key
	minPQ.swim(minPQ.qp[i])
	minPQ.sink(minPQ.qp[i])
}

func (minPQ IndexMinPQ) less(i, j int) bool {
	cmp := minPQ.keys[minPQ.pq[i]].CompareTo(minPQ.keys[minPQ.pq[j]])
	return cmp < 0
}

// swap invariance: qp[pq[i]] = pq[qp[i]] = i
func (minPQ *IndexMinPQ) swap(i, j int) {
	minPQ.qp[minPQ.pq[i]], minPQ.qp[minPQ.pq[j]] = minPQ.qp[minPQ.pq[j]], minPQ.qp[minPQ.pq[i]]
	minPQ.pq[i], minPQ.pq[j] = minPQ.pq[j], minPQ.pq[i]
	//minPQ.qp[minPQ.pq[i]] = i
	//minPQ.qp[minPQ.pq[j]] = j
}

func (minPQ *IndexMinPQ) swim(k int) {
	for ; k > 1; k = k / 2 {
		if minPQ.less(k, k/2) {
			minPQ.swap(k, k/2)
		} else {
			break
		}
	}
}

func (minPQ *IndexMinPQ) sink(k int) {
	for 2*k <= minPQ.n {
		j := 2 * k
		if j < minPQ.n && minPQ.less(j+1, j) {
			j++
		}
		if minPQ.less(j, k) {
			minPQ.swap(j, k)
		} else {
			break
		}
		k = j
	}
}

func (minPQ IndexMinPQ) MinIndex() int {
	if minPQ.n == 0 {
		fmt.Println("MinIndex error: priority queue underflow")
		return -1
	}
	return minPQ.pq[1]
}

func (minPQ IndexMinPQ) MinKey() typ.Comparable {
	if minPQ.n == 0 {
		fmt.Println("MinKey error: priority queue underflow")
		return nil
	}
	return minPQ.keys[minPQ.pq[1]]
}

// DelMin deletes the minimum key and returns the corresponding index.
func (minPQ *IndexMinPQ) DelMin() int {
	if minPQ.n == 0 {
		fmt.Println("DelMin error: priority queue underflow")
		return -1
	}
	min := minPQ.pq[1]
	minPQ.swap(1, minPQ.n)
	minPQ.n--
	minPQ.sink(1)
	minPQ.qp[min] = -1
	minPQ.keys[min] = nil
	minPQ.pq[minPQ.n+1] = -1
	return min
}

func (minPQ IndexMinPQ) KeyOf(i int) typ.Comparable {
	err := minPQ.validateIndex(i)
	if err != nil {
		fmt.Println("KeyOf error:", err)
		return nil
	}
	return minPQ.keys[i]
}
