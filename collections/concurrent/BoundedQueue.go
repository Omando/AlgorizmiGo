package concurrent

import (
	"container/list"
	"fmt"
	"sync"
)

type BoundedBuffer struct {
	Capacity  int
	Data      *list.List
	Condition *sync.Cond
}

func (bb *BoundedBuffer) isFull() bool {
	//return len(bb.Data) == cap(bb.Data)
	return bb.Data.Len() == bb.Capacity
}

func (bb *BoundedBuffer) isEmpty() bool {
	return bb.Data.Len() == 0
}

func (bb *BoundedBuffer) Enqueue(item int) {
	bb.Condition.L.Lock()
	defer bb.Condition.L.Unlock()
	for bb.isFull() {
		fmt.Println("Cannot enqueue. Buffer full. Waiting to dequeue")
		bb.Condition.Wait()
	}

	fmt.Println("Buffer not full. Enqueueing")
	bb.Data.PushBack(item)
	bb.Condition.Broadcast()
}

func (bb *BoundedBuffer) Dequeue() int {
	bb.Condition.L.Lock()
	defer bb.Condition.L.Unlock()
	for bb.isEmpty() {
		fmt.Println("Cannot dequeue. Buffer empty. Waiting to enqueue")
		bb.Condition.Wait()
	}

	// Get item
	fmt.Println("Buffer not empty. Dequeuing")
	front := bb.Data.Front()
	bb.Data.Remove(front)

	// Wake all goroutines
	bb.Condition.Broadcast()

	return front.Value.(int)
}

func (bb *BoundedBuffer) Error() string {
	return fmt.Sprintf("Cannot remove. List is empty")
}
