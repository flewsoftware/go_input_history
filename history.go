package input_history

import (
	"sync"
)

// History struct to hold the history of the user input
type History struct {
	Max   int
	Data  []string
	Index int

	mu sync.Mutex
}

// New history struct
func New(max int) *History {
	return &History{
		Max:   max,
		Data:  []string{},
		Index: 0,
	}
}

// Add line to history
func (h *History) Add(line string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	previous := h.Prev()

	if line == previous {
		return
	}

	// remove the left most element if we have reached the max to make way for the new element
	if len(h.Data) > h.Max {
		h.Data = append(h.Data[:0], h.Data[1:]...)
	}

	h.Data = append(h.Data, line)
	h.Index = len(h.Data) - 1
}

// HasLine will return true if the history buffer contains the given string
func (h *History) HasLine(line string) bool {
	for _, l := range h.Data {
		if line == l {
			return true
		}
	}

	return false
}

// Get history at index
func (h *History) Get(index int) string {
	return h.Data[index]
}

// IsEmpty return true if the history data array is empty
func (h *History) IsEmpty() bool {
	if len(h.Data) <= 0 {
		return true
	}

	return false
}

// Prev returns the previous history line from the current index
func (h *History) Prev() string {

	h.Index--

	if h.Index < 0 {
		h.Index = 0
	}

	if h.IsEmpty() {
		return ""
	}

	return h.Data[h.Index]
}

// Next returns the history line after the current index
func (h *History) Next() string {

	h.Index++

	if h.Index >= len(h.Data) {
		h.Index = len(h.Data) - 1
	}

	if h.IsEmpty() {
		return ""
	}

	return h.Data[h.Index]
}

// SetIndexToNew sets the index to the new elements index
func (h *History) SetIndexToNew() {
	h.Index = len(h.Data)
}
