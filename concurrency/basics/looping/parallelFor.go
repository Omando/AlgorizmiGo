package looping

import "sync"

type ActionFunc = func(data interface{}) interface{}

// ParallelFor Starting implementation of a parallel for with a given function type
// This initial implementation is a place holder and will be enhanced/modified
// as required
// Enhancement 1: provide updates to client via a channel
// Enhancement 2: parameterize the return type
func ParallelFor(data []interface{}, action ActionFunc) (results []interface{}) {
	var wg sync.WaitGroup
	var lock sync.Mutex
	wg.Add(len(data))
	for _,datum := range data {
		go func(data interface{}) {
			defer wg.Done()
			result := action(data)		// some expensive operation
			lock.Lock()
			results = append(results, result )
			lock.Unlock()
		}(datum)
	}

	wg.Wait()

	return results
}
