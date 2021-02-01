package stack

import "sync"

// Item description
type Item interface {}

// Stack description
type Stack struct {
	Items []Item
	lock sync.Mutex
}

// Push description
func (st *Stack) Push(elem Item) {
	st.lock.Lock()
	defer st.lock.Unlock()

	st.Items = append(st.Items, elem)
}


// Pop description
func (st *Stack) Pop() Item {
	st.lock.Lock()
	defer st.lock.Unlock()

	if len(st.Items) == 0 {
		return nil
	}

	l := len(st.Items)

	oldItem := st.Items[l-1]
	st.Items = st.Items[:l-1]

	return oldItem
}


// IsEmpty description
func (st *Stack) IsEmpty() bool {
	st.lock.Lock()
	defer st.lock.Unlock()

	return len(st.Items) == 0
}


// Size description
func (st *Stack) Size() int {
	st.lock.Lock()
	defer st.lock.Unlock()

	return len(st.Items)
}


// Clear description
func (st *Stack) Clear() {
	st.lock.Lock()
	defer st.lock.Unlock()

	st.Items = nil
}


// Dump description
func (st *Stack) Dump() []Item {
	st.lock.Lock()
	defer st.lock.Unlock()

	tmpItems := make([]Item, len(st.Items))
	copy(tmpItems, st.Items)
	return tmpItems
}


// Peek description
func (st *Stack) Peek() Item {
	st.lock.Lock()
	defer st.lock.Unlock()

	if len(st.Items) == 0 {
		return nil
	}

	return st.Items[len(st.Items)-1]
}


