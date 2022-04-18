package promise

import (
	"context"
)

// Race return only once result who is the first done.
func Race(promises ...*Promise) *Promise {
	return New(func(resolve Resolver, reject Rejecter) {
		ch := make(chan struct{})
		ctx, cancel := context.WithCancel(context.Background())

		var result interface{}
		var err error

		defer func() {
			if err != nil {
				reject(err)
			} else {
				resolve(result)
			}
		}()

		for _, p := range promises {
			go func(p *Promise) {
				select {
				// case <-p.Wait():
				case <-p.ch:
					r, e := p.result, p.err
					if e != nil {
						err = e
					} else {
						result = r
					}

					ch <- struct{}{}
				case <-ctx.Done():
					return
				}
			}(p)
		}

		<-ch
		cancel()
	})
}
