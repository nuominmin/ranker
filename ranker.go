package ranker

import (
	"container/heap"
	"sort"
)

type item[T any] struct {
	Value T
	Score int64 // 排序依据的分数
}

type Ranker[T any] struct {
	topN      int
	priorityQ PriorityQueue[T]
}

func NewRanker[T any](topN int) *Ranker[T] {
	return &Ranker[T]{
		topN:      topN,
		priorityQ: make(PriorityQueue[T], 0, topN),
	}
}

// Add 添加数据，并根据分数进行排序
func (r *Ranker[T]) Add(score int64, value T) {
	if len(r.priorityQ) < r.topN {
		heap.Push(&r.priorityQ, item[T]{
			Value: value,
			Score: score,
		})
	} else if score > r.priorityQ[0].Score {
		heap.Pop(&r.priorityQ)
		heap.Push(&r.priorityQ, item[T]{
			Value: value,
			Score: score,
		})
	}
}

func (r *Ranker[T]) GetTopN() []T {
	// 创建副本，避免改变原有堆的结构
	pqCopy := make(PriorityQueue[T], len(r.priorityQ))
	copy(pqCopy, r.priorityQ)

	// 对副本进行排序
	sort.Slice(pqCopy, func(i, j int) bool {
		return pqCopy[i].Score > pqCopy[j].Score
	})

	// 提取排序后的值
	sortedItems := make([]T, len(pqCopy))
	for i := 0; i < len(pqCopy); i++ {
		sortedItems[i] = pqCopy[i].Value
	}
	return sortedItems
}

// PriorityQueue 实现最小堆，用于存储和排序数据
type PriorityQueue[T any] []item[T]

func (pq *PriorityQueue[T]) Len() int { return len(*pq) }

func (pq *PriorityQueue[T]) Less(i, j int) bool {
	return (*pq)[i].Score < (*pq)[j].Score
}

func (pq *PriorityQueue[T]) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue[T]) Push(x interface{}) {
	*pq = append(*pq, x.(item[T]))
}

func (pq *PriorityQueue[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	data := old[n-1]
	*pq = old[0 : n-1]
	return data
}
