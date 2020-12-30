package concurrency_basics

import "sync"

type ActionFunc = func(data interface{}) interface{}

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
