package promise

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestPromise(t *testing.T) {
	results := []string{}
	var promiseErr error

	p := New(func(resolve Resolver, reject Rejecter) {
		time.Sleep(100 * time.Millisecond)
		results = append(results, "0")

		resolve(1)
	})

	p.Then(func(v interface{}) interface{} {
		results = append(results, fmt.Sprintf("%d", v.(int)))
		return v.(int) + 1
	})

	p.Then(func(v interface{}) interface{} {
		results = append(results, fmt.Sprintf("%d", v.(int)))
		return v.(int) * 2
	})

	p.Catch(func(err error) {
		promiseErr = err
	})

	p.Then(func(v interface{}) interface{} {
		results = append(results, fmt.Sprintf("%d", v.(int)))
		return v.(int) + 3
	})

	res, err := p.Wait()
	if err != nil {
		t.Error(err)
	}

	results = append(results, "done")

	if promiseErr != nil {
		t.Errorf("expected no error, got %v", promiseErr)
	}

	if strings.Join(results, ",") != "0,1,2,4,done" {
		t.Errorf("expected '0,1,2,4,done', got %s", strings.Join(results, ","))
	}

	if res != 7 {
		t.Errorf("expected 7, got %v", p.result)
	}
}

func TestPromiseCatch(t *testing.T) {
	results := []string{}
	var promiseErr error

	p := New(func(resolve Resolver, reject Rejecter) {
		time.Sleep(100 * time.Millisecond)
		results = append(results, "0")

		panic("error occurred")
	})

	p.Then(func(v interface{}) interface{} {
		results = append(results, fmt.Sprintf("%d", v.(int)))
		return v.(int) + 1
	})

	p.Then(func(v interface{}) interface{} {
		results = append(results, fmt.Sprintf("%d", v.(int)))
		return v.(int) * 2
	})

	p.Catch(func(err error) {
		promiseErr = err
	})

	p.Then(func(v interface{}) interface{} {
		results = append(results, fmt.Sprintf("%d", v.(int)))
		return v.(int) + 3
	})

	res, err := p.Wait()
	if err == nil {
		t.Error(err)
	}

	results = append(results, "done")

	if promiseErr == nil {
		t.Errorf("expected no error, got %v", promiseErr)
	}

	if strings.Join(results, ",") != "0,done" {
		t.Errorf("expected '0,done', got %s", strings.Join(results, ","))
	}

	if res != nil {
		t.Errorf("expected nil, got %v", p.result)
	}
}
