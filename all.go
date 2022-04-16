package promise

import (
	"sync"
)

// All handles all promises.
func All(promises ...*Promise) *Promise {
	return New(func(resolve Resolver, reject Rejecter) {
		var results []interface{}
		var err error

		wg := &sync.WaitGroup{}

		defer func() {
			if err != nil {
				reject(err)
			} else {
				resolve(results)
			}
		}()

		for _, p := range promises {
			p.Then(func(v interface{}) interface{} {
				results = append(results, v)
				return v
			}).Catch(func(e error) {
				err = e
			}).Finally(func() {
				defer wg.Done()
			})

			wg.Add(1)
		}

		wg.Wait()
	})
}
