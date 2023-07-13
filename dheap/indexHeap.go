package dheap

type IndexHeap[T Item] struct {
	len        int
	cap        int
	items      []T
	indexes    []int
	verIndexes []int
}

func NewIndexHeap[T Item](cap int) *IndexHeap[T] {
	return &IndexHeap[T]{
		cap:        cap,
		items:      make([]T, cap+1),
		indexes:    make([]int, cap+1),
		verIndexes: make([]int, cap+1),
	}
}

// shift up
func (h *IndexHeap[T]) Push(id int, item T) {
	//防止越界
	if id+1 > h.cap {
		return
	}
	//判断扩容
	if h.len == h.cap {
		h.extend()
	}
	id = id + 1
	h.items[id] = item

	h.len++
	p := h.len
	h.indexes[p] = id
	h.verIndexes[h.indexes[p]] = p
	h.shiftUp(p)

}

// shift down
func (h *IndexHeap[T]) Pop() (T, int) {
	ret := h.items[h.indexes[1]]
	idx := h.indexes[1] - 1
	h.indexes[1], h.indexes[h.len] = h.indexes[h.len], h.indexes[1]
	h.verIndexes[h.indexes[1]] = 1
	h.verIndexes[h.indexes[h.len]] = h.len
	h.len--
	p := 1
	h.shiftDown(p)
	return ret, idx
}
func (h *IndexHeap[T]) GetIndex() int {
	return h.indexes[1] - 1
}
func (h *IndexHeap[T]) GetItemByIndex(idx int) T {
	return h.items[idx+1]
}
func (h *IndexHeap[T]) IsContain(idx int) bool {
	return h.items[idx+1] != nil
}
func (h *IndexHeap[T]) Change(id int, modifiedItem T) {
	id = id + 1
	h.items[id] = modifiedItem

	p := h.verIndexes[id]
	h.shiftDown(p)
	h.shiftUp(p)

}
func (h *IndexHeap[T]) Len() int {
	return h.len
}

func (h *IndexHeap[T]) extend() {
	//fmt.Printf("发生扩容，从%v到%v\n", h.cap, 2*h.cap)
	d := h.cap + 1
	h.cap = 2 * h.cap
	extend_items := make([]T, h.cap+1)
	extend_indexes := make([]int, h.cap+1)
	extend_verIndexes := make([]int, h.cap+1)
	copy(extend_items[1:d], h.items[1:])
	copy(extend_indexes[1:d], h.indexes[1:])
	copy(extend_verIndexes[1:d], h.verIndexes[1:])
	h.items = extend_items
	h.indexes = extend_indexes
	h.verIndexes = extend_verIndexes

}

func (h *IndexHeap[T]) shiftDown(p int) {

	for p*2 <= h.len {
		j := 2 * p
		if j+1 <= h.len && h.items[h.indexes[j+1]].Greater(h.items[h.indexes[j]]) {
			j = j + 1
		}
		if h.items[h.indexes[p]].Greater(h.items[h.indexes[j]]) {
			break
		}
		h.indexes[p], h.indexes[j] = h.indexes[j], h.indexes[p]
		h.verIndexes[h.indexes[p]] = p
		h.verIndexes[h.indexes[j]] = j
		p = j
	}
}

func (h *IndexHeap[T]) shiftUp(p int) {

	item := h.items[h.indexes[p]]
	for p > 1 && item.Greater(h.items[h.indexes[p/2]]) {
		h.indexes[p], h.indexes[p/2] = h.indexes[p/2], h.indexes[p]
		h.verIndexes[h.indexes[p]] = p
		h.verIndexes[h.indexes[p/2]] = p / 2
		p = p / 2
	}
}
