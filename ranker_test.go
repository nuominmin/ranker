package ranker

import (
	"container/heap"
	"github.com/shopspring/decimal"
	"testing"
)

func TestRanker_AddAndGetTopN(t *testing.T) {
	r := NewRanker[int](3)

	// 添加数据
	r.Add(decimal.NewFromInt(10), 1)
	r.Add(decimal.NewFromInt(20), 2)
	r.Add(decimal.NewFromInt(15), 3)
	r.Add(decimal.NewFromInt(25), 4)
	r.Add(decimal.NewFromInt(5), 5)

	// 获取前N名
	topItems := r.GetTopN()

	// 预期结果应该是分数最高的三个元素
	expected := []int{4, 2, 3}

	if len(topItems) != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), len(topItems))
	}

	for i := range expected {
		if topItems[i] != expected[i] {
			t.Errorf("expected item %d, got %d", expected[i], topItems[i])
		}
	}
}

func TestRanker_AddLessThanTopN(t *testing.T) {
	r := NewRanker[int](3)

	// 添加少量数据
	r.Add(decimal.NewFromInt(10), 1)
	r.Add(decimal.NewFromInt(20), 2)

	// 获取前N名
	topItems := r.GetTopN()

	// 预期结果应该是所有添加的元素
	expected := []int{2, 1}

	if len(topItems) != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), len(topItems))
	}

	for i := range expected {
		if topItems[i] != expected[i] {
			t.Errorf("expected item %d, got %d", expected[i], topItems[i])
		}
	}
}

func TestRanker_Empty(t *testing.T) {
	r := NewRanker[int](3)

	// 获取前N名
	topItems := r.GetTopN()

	// 预期结果应该为空
	if len(topItems) != 0 {
		t.Fatalf("expected length 0, got %d", len(topItems))
	}
}

func TestPriorityQueue(t *testing.T) {
	pq := make(PriorityQueue[int], 0)
	heap.Init(&pq)

	heap.Push(&pq, item[int]{Value: 1, Score: decimal.NewFromInt(10)})
	heap.Push(&pq, item[int]{Value: 2, Score: decimal.NewFromInt(20)})
	heap.Push(&pq, item[int]{Value: 3, Score: decimal.NewFromInt(5)})

	if pq.Len() != 3 {
		t.Fatalf("expected length 3, got %d", pq.Len())
	}

	minItem := heap.Pop(&pq).(item[int])

	// 由于这是一个最小堆，最小的元素应该是 Score = 5 的那个
	if !minItem.Score.Equal(decimal.NewFromInt(5)) {
		t.Errorf("expected minimum score 5, got %d", minItem.Score)
	}
}
