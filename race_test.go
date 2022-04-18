package promise

import (
	"testing"
	"time"
)

func TestRace(t *testing.T) {
	p := Race(
		New(func(resolve Resolver, reject Rejecter) {
			// fmt.Println("resolve start 1")
			time.Sleep(200 * time.Millisecond)
			// fmt.Println("resolve 1")
			resolve(1)
		}),
		New(func(resolve Resolver, reject Rejecter) {
			// fmt.Println("resolve start 2")
			time.Sleep(100 * time.Millisecond)
			// fmt.Println("resolve 2")
			resolve(2)
		}),
		New(func(resolve Resolver, reject Rejecter) {
			// fmt.Println("resolve start 3")
			time.Sleep(300 * time.Millisecond)
			// fmt.Println("resolve 3")
			resolve(3)
		}),
	)

	resultX, err := p.Wait()
	if err != nil {
		t.Error(err)
	}

	time.Sleep(400 * time.Millisecond)

	result := resultX.(int)
	if result != 2 {
		t.Errorf("expected 2, got %v", result)
	}
}
