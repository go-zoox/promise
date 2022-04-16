package promise

// Resolve return a promise that resolve with the given value.
func Resolve(v interface{}) *Promise {
	return New(func(resolve Resolver, reject Rejecter) {
		resolve(v)
	})
}

// Reject return a promise that reject with the given error.
func Reject(err error) *Promise {
	return New(func(resolve Resolver, reject Rejecter) {
		reject(err)
	})
}
