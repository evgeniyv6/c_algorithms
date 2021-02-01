package queue

import "sync"


// Item description
type Item interface {}

// Queue description
type Queue struct {
	items []Item
	lock sync.Mutex
}

// exercise 5: queue with array

// ArrEnqueue description
func (q *Queue) ArrEnqueue(elem Item) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.items = append(q.items, elem)
}

// ArrDequeue description
func (q *Queue) ArrDequeue() Item {
	q.lock.Lock()
	defer q.lock.Unlock()

	if len(q.items) == 0 {
		return nil
	}
	oldItem := q.items[0]
	q.items = q.items[1:]
	return oldItem
}


// Dump description
func (q *Queue) Dump() []Item {
	q.lock.Lock()
	defer q.lock.Unlock()

	tmpItems := make([]Item, len(q.items))
	copy(tmpItems, q.items)
	return tmpItems
}
