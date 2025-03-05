package semgroup

import (
	"fmt"
	"sync"
)

type token struct{}

// A Group is a collection of goroutines working on subtasks that are part of
// the same overall task.
//
// A zero Group is valid and has no limit on the number of active goroutines.
type Group struct {
	wg  sync.WaitGroup
	sem chan token
}

func (g *Group) done() {
	if g.sem != nil {
		<-g.sem
	}
	g.wg.Done()
}

// Wait blocks until all function calls from the Go method have returned.
func (g *Group) Wait() {
	g.wg.Wait()
}

// Go calls the given function in a new goroutine.
// It blocks until the new goroutine can be added without the number of
// active goroutines in the group exceeding the configured limit.
func (g *Group) Go(f func()) {
	if g.sem != nil {
		g.sem <- token{}
	}

	g.wg.Add(1)
	go func() {
		defer g.done()
		f()
	}()
}

// SetLimit limits the number of active goroutines in this group to at most n.
// A negative value indicates no limit.
//
// Any subsequent call to the Go method will block until it can add an active
// goroutine without exceeding the configured limit.
//
// The limit must not be modified while any goroutines in the group are active.
func (g *Group) SetLimit(n int) {
	if n < 0 {
		g.sem = nil
		return
	}
	if len(g.sem) != 0 {
		panic(fmt.Errorf("semgroup: modify limit while %v goroutines in the group are still active", len(g.sem)))
	}
	g.sem = make(chan token, n)
}
