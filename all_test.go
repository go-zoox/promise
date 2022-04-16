package promise

import (
	"testing"
	"time"
)

func TestAll(t *testing.T) {
	ps := All(
		New(func(resolve Resolver, reject Rejecter) {
			time.Sleep(100 * time.Millisecond)
			resolve(1)
		}),
		New(func(resolve Resolver, reject Rejecter) {
			time.Sleep(10 * time.Millisecond)
			resolve(2)
		}),
		New(func(resolve Resolver, reject Rejecter) {
			time.Sleep(110 * time.Millisecond)
			resolve(3)
		}),
	)

	resultsX, err := ps.Wait()
	if err != nil {
		t.Error(err)
	}

	results := resultsX.([]interface{})

	if len(results) != 3 {
		t.Errorf("expected 3 results, got %d", len(results))
	}

	if results[0].(int) != 2 {
		t.Errorf("expected 2, got %v", results[0])
	}

	if results[1].(int) != 1 {
		t.Errorf("expected 1, got %v", results[1])
	}

	if results[2].(int) != 3 {
		t.Errorf("expected 3, got %v", results[2])
	}
}
