package promise

// // Race return only once result who is the first done.
// func Race(promises ...*Promise) *Promise {
// 	return New(func(resolve Resolver, reject Rejecter) {
// 		ctx, cancel := context.WithCancel(context.Background())

// 		var result interface{}
// 		var err error

// 		defer func() {
// 			if err != nil {
// 				reject(err)
// 			} else {
// 				resolve(result)
// 			}
// 		}()

// 		for _, p := range promises {
// 			go func(p *Promise, cancel context.CancelFunc) {
// 				// p.Then(func(v interface{}) interface{} {
// 				// 	return v
// 				// }).Catch(func(e error) {
// 				// 	err = e
// 				// })

// 				// r, e := p.Wait()
// 				// if e != nil {
// 				// 	err = e
// 				// } else {
// 				// 	result = r
// 				// }

// 				fmt.Println("xxx")
// 				cancel()
// 			}(p, cancel)
// 		}

// 		<-ctx.Done()
// 		cancel()
// 	})
// }
