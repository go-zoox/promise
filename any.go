package promise

// Any returns a promise that resolve with any one promise done.
func Any(promises ...*Promise) *Promise {
	return Race(promises...)
}
