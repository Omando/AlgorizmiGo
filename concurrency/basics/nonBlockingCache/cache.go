package nonBlockingCache

import "sync"

// result encapsulates the return value of httpGetBody
type result struct {
	value interface{}
	err error
}

type entry struct {
	res result
	ready chan struct{}		// closed when res is ready
}

// Func is the type of the function whose result will be memoized
type Func func(key string) (interface{}, error)

// Memo caches the results of calling a Func
type Memo struct {
	f Func
	cache map[string]*entry
	lock sync.Mutex
}

// Create an empty cache for a Func
// A Memo instance holds the function f to memoize, of type Func, and the cache
// which is a  mapping from strings to results
func New(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]*entry),
	}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.lock.Lock()

	// Try to get the key's value
	var e *entry = memo.cache[key]
	if  e == nil {
		// Data is not available. Create a cache entry for this key, but keep the entry's
		// channel open until the result is ready. Other clients can get this entry, but
		// the channel will block until data is ready
		e = &entry{
			ready: make(chan struct{}),
		}
		memo.cache[key] = e
		memo.lock.Unlock()

		// Perform initialization. Only one goroutine can ever execute this line
		e.res.value, e.res.err = memo.f(key);

		// Data is now ready. Broadcast over the channel
		close(e.ready)
	} else {
		// Entry is available so unlock. But...
		memo.lock.Unlock()

		// data may not be ready yet so wait for it by blocking (if required) on the entry's channel
		<- e.ready
	}
	return e.res.value, e.res.err
}