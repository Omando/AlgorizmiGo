package nonBlockingCache

// result encapsulates the return value of httpGetBody
type result struct {
	value interface{}
	err error
}

// Func is the type of the function whose result will be memoized
type Func func(key string) (interface{}, error)

// Memo caches the results of calling a Func
type Memo struct {
	f Func
	cache map[string]result
}

// Create an empty cache for a Func
// A Memo instance holds the function f to memoize, of type Func, and the cache
// which is a  mapping from strings to results
func New(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]result),
	}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	result, ok := memo.cache[key]
	if  !ok {
		// Invoke the function lazily and cache its value
		result.value, result.err = memo.f(key);
		memo.cache[key] = result
	}
	return result.value, result.err
}