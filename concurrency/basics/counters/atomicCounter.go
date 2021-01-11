package counters

import "sync"

type AtomicCounter struct {
	count int
	lock sync.Mutex
}

func (ac *AtomicCounter) Increment() {
	ac.lock.Lock()
	ac.count += 1
	ac.lock.Unlock()
}

func (ac *AtomicCounter) Get() int {
	return ac.count
}


