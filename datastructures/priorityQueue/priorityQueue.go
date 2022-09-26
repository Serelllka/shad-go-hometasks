package priorityQueue

type item[T any] struct {
	value    T
	priority int
	index    int
}

type PriorityQueue[T any] []*item[T]

func New[T any](cap int) PriorityQueue[T] {
	return make([]*item[T], 0, cap)
}

func (pq PriorityQueue[T]) Len() int {
	return len(pq)
}

func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*item[T])
	item.index = n
	*pq = append(*pq, item)
}
