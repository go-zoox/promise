package promise

import (
	"errors"
	"fmt"
)

// Promise is promise object.
type Promise struct {
	ch     chan struct{}
	cb     func(resolve Resolver, reject Rejecter)
	result interface{}
	err    error
	//
	thens   []Thener
	catch   Catcher
	finally func()
}

type Resolver func(v interface{})
type Rejecter func(err error)

type Thener func(v interface{}) interface{}
type Catcher func(err error)

// New creates a new promise.
func New(cb func(resolve Resolver, reject Rejecter)) *Promise {
	p := &Promise{
		ch: make(chan struct{}),
		cb: cb,
	}

	p.start()

	return p
}

func (p *Promise) start() {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				switch v := err.(type) {
				case error:
					p.err = v
				case string:
					p.err = errors.New(v)
				default:
					p.err = fmt.Errorf("%v", v)
				}
				p.executeRejector()
			}

			if p.finally != nil {
				p.finally()
			}
			p.ch <- struct{}{}
		}()
		p.cb(func(v interface{}) {
			p.result = v
			p.executeResolvers()
		}, func(err error) {
			p.err = err
			p.executeRejector()
		})
	}()
}

func (p *Promise) executeResolvers() {
	for _, resolver := range p.thens {
		p.result = resolver(p.result)
	}
}

func (p *Promise) executeRejector() {
	p.catch(p.err)
}

// Then adds a new then handler.
func (p *Promise) Then(resolve Thener) *Promise {
	p.thens = append(p.thens, resolve)
	return p
}

// Catch adds a new catch handler.
func (p *Promise) Catch(reject Catcher) *Promise {
	p.catch = reject
	return p
}

// Finally adds a new finally handler.
func (p *Promise) Finally(f func()) *Promise {
	p.finally = f
	return p
}

// Wait waits for the promise to be resolved or rejected.
func (p *Promise) Wait() (interface{}, error) {
	<-p.ch
	return p.result, p.err
}
