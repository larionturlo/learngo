package heap

const (
	minHeap HeapType = -1
	maxHeap HeapType = 1
)

type HeapType int

type Heap struct {
	heap     []int
	heapType HeapType
}

func NewHeap(t HeapType) *Heap {
	return &Heap{
		heap:     []int{},
		heapType: t,
	}
}

func parent(i int) int {
	return i / 2
}
func left(i int) int {
	return i * 2
}
func right(i int) int {
	return i*2 + 1
}
func comparator(child, parent int, heapType HeapType) HeapType {
	if child < parent {
		return minHeap
	}
	if child > parent {
		return maxHeap
	}
	return heapType
}

func (h *Heap) bubbleUp(i int) {
	if i >= len(h.heap) {
		return
	}
	p := parent(i)

	comp := comparator(h.heap[p], h.heap[i], h.heapType)

	if h.heapType == comp {
		return
	}

	h.heap[i], h.heap[p] = h.heap[p], h.heap[i]

	h.bubbleUp(p)
}

func (h *Heap) bubbleDown(i int) {
	if i >= len(h.heap) {
		return
	}
	l := left(i)
	r := right(i)

	child := l
	if child < len(h.heap) && r < len(h.heap) {
		compch := comparator(h.heap[child], h.heap[r], h.heapType)

		if h.heapType != compch {
			child = r
		}
	} else if child >= len(h.heap) {
		if r >= len(h.heap) {
			return
		}
		child = r
	}

	comp := comparator(h.heap[i], h.heap[child], h.heapType)

	if h.heapType == comp {
		return
	}

	h.heap[i], h.heap[child] = h.heap[child], h.heap[i]

	h.bubbleDown(child)
}

// Function return heap as slice
func (h *Heap) AsSlice() []int {
	return h.heap
}

// Function for heapify slice
func (h *Heap) Heapify(arr []int) {
	h.heap = arr
	for i := len(arr) - 1; i >= 0; i-- {
		h.bubbleDown(i)
	}
}

// Function insert a number to the heap
func (h *Heap) Insert(n int) {
	h.heap = append(h.heap, n)
	h.bubbleUp(len(h.heap) - 1)
}

// Function delete a number to the heap
func (h *Heap) Extract(i int) int {
	n := 0
	last := len(h.heap) - 1
	if i < 0 || i > last {
		return n
	}
	h.heap[i], h.heap[last] = h.heap[last], h.heap[i]
	n = h.heap[last]
	h.heap = h.heap[:last]
	h.bubbleDown(i)
	return n
}
