package history

import "sync"

// Simple is a History that only has a single stack
type Simple struct {
	items []*Item
	mu    sync.RWMutex
}

func NewSimple() *Simple {
	return &Simple{}
}

func (h *Simple) Push(item *Item) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.items = append(h.items, item)
}

func (h *Simple) Pop() *Item {
	h.mu.Lock()
	defer h.mu.Unlock()

	if len(h.items) == 0 {
		return nil
	}

	item := h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]

	return item
}

func (h *Simple) Peek() *Item {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.items[len(h.items)-1]
}

func (h *Simple) Depth() int {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return len(h.items)
}

func (h *Simple) All() []*Item {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.items
}

func (h *Simple) Clear() {
	h.mu.Lock()
	defer h.mu.Unlock()

	// todo: finishers?
	h.items = make([]*Item, 0)
}
