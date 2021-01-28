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

}