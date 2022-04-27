package erronce

import "sync"

// ErrOnce is like sync.Once, but the Do() accepts a func() error and can
// therefore also return an error.
type ErrOnce struct {
	once sync.Once
	err  error
}

func (o *ErrOnce) Do(fn func() error) error {
	o.once.Do(func() {
		o.err = fn()
	})
	return o.err
}
