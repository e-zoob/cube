package queue

type (
	Queue struct {
		start, end *node
		lenght     int
	}
	node struct {
		value interface{}
		next  *node
	}
)

// Create a new queue
func New() *Queue {
	return &Queue{nil, nil, 0}

}

// Take the next item off the front of the queue
func (this *Queue) Dequeue() interface{} {
	if this.lenght == 0 {
		return nil
	}

	n := this.start
	if this.lenght == 1 {
		this.start = nil
		this.end = nil
	} else {
		this.start = this.start.next
	}
	this.lenght--
	return n.value
}

// put an item at the end of a queue
func (this *Queue) Enqueue(value interface{}) {
	n := &node{value, nil}
	if this.lenght == 0 {
		this.start = n
		this.end = n
	} else {
		this.end.next = n
		this.end = n
	}
	this.lenght++
}

// Return the numern of the items in the queue
func (this *Queue) Len() int {
	return this.lenght
}

// return the first item in the queue without removing it
func (this *Queue) Peek() interface{} {
	if this.lenght == 0 {
		return nil
	}
	return this.start.value
}
