package ranker

import (
	"container/heap"
	"github.com/shopspring/decimal"
	"sort"
)

type item[T any] struct {
	value T
	score decimal.Decimal // 排序依据的分数
}

type Ranker[T any] struct {
	topN      int
	priorityQ priorityQueue[T]
}

func NewRanker[T any](topN int) *Ranker[T] {
	return &Ranker[T]{
		topN:      topN,
		priorityQ: make(priorityQueue[T], 0, topN),
	}
}

// Add 添加数据，并根据分数进行排序
func (r *Ranker[T]) Add(score decimal.Decimal, value T) {
	if len(r.priorityQ) < r.topN {
		heap.Push(&r.priorityQ, item[T]{
			value: value,
			score: score,
		})
		return
	}

	if score.GreaterThan(r.priorityQ[0].score) {
		heap.Pop(&r.priorityQ)
		heap.Push(&r.priorityQ, item[T]{
			value: value,
			score: score,
		})
	}
}

func (r *Ranker[T]) GetTopN() []T {
	// 创建副本，避免改变原有堆的结构
	pqCopy := make(priorityQueue[T], len(r.priorityQ))
	copy(pqCopy, r.priorityQ)

	// 对副本进行排序
	sort.Slice(pqCopy, func(i, j int) bool {
		return pqCopy[i].score.GreaterThan(pqCopy[j].score)
	})

	// 提取排序后的值
	sortedItems := make([]T, len(pqCopy))
	for i := 0; i < len(pqCopy); i++ {
		sortedItems[i] = pqCopy[i].value
	}
	return sortedItems
}

// priorityQueue 实现最小堆，用于存储和排序数据
type priorityQueue[T any] []item[T]

func (pq *priorityQueue[T]) Len() int { return len(*pq) }

func (pq *priorityQueue[T]) Less(i, j int) bool {
	return (*pq)[i].score.LessThan((*pq)[j].score)
}

func (pq *priorityQueue[T]) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *priorityQueue[T]) Push(x interface{}) {
	*pq = append(*pq, x.(item[T]))
}

func (pq *priorityQueue[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	data := old[n-1]
	*pq = old[0 : n-1]
	return data
}
