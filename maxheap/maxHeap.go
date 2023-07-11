package maxheap

type Item interface {
	Greater(item Item) bool
}
type MaxHeap[T Item] struct {
	len   int
	cap   int
	items []T
}

func NewMaxHeap[T Item](cap int) *MaxHeap[T] {
	return &MaxHeap[T]{
		cap:   cap,
		items: make([]T, cap+1),
	}
}

// shift up
func (h *MaxHeap[T]) Push(item T) {
	if h.len == h.cap {
		h.Extend()
	}
	h.len++
	h.items[h.len] = item
	p := h.len
	for p > 1 && item.Greater(h.items[p/2]) {
		h.items[p], h.items[p/2] = h.items[p/2], h.items[p]
		p = p / 2
	}
}

// shift down
func (h *MaxHeap[T]) Pop() T {
	ret := h.items[1]

	h.items[1], h.items[h.len] = h.items[h.len], h.items[1]
	h.len--
	p := 1
	for p*2 <= h.len {
		j := 2 * p
		if j+1 <= h.len && h.items[j+1].Greater(h.items[j]) {
			j = j + 1
		}
		if h.items[p].Greater(h.items[j]) {
			break
		}
		h.items[p], h.items[j] = h.items[j], h.items[p]
		p = j
	}
	return ret
}
func (h *MaxHeap[T]) Len() int {
	return h.len
}

func (h *MaxHeap[T]) Extend() {
	//fmt.Printf("发生扩容，从%v到%v\n", h.cap, 2*h.cap)
	d := h.cap + 1
	h.cap = 2 * h.cap
	extend := make([]T, h.cap+1)
	copy(extend[1:d], h.items[1:])
	h.items = extend

}
