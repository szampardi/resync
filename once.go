package resync

import (
	"sync"
	"sync/atomic"
)

// Once is an object that will perform exactly one action
// until Reset is called.
// See http://golang.org/pkg/sync/#Once
type Once struct {
	Fn   func() error
	m    sync.Mutex
	done uint32
	err  error
}

// Do simulates sync.Once.Do by executing the specified function
// only once, until Reset is called.
// See http://golang.org/pkg/sync/#Once
func (o *Once) Do() {
	switch atomic.LoadUint32(&o.done) {
	case 1:
		return
	case 0:
		o.m.Lock()
		defer o.m.Unlock()
		defer atomic.StoreUint32(&o.done, 1)
		o.err = o.Fn()
	}
}

// Reset indicates that the next call to Do should actually be called
// once again.
func (o *Once) Reset() {
	o.m.Lock()
	defer o.m.Unlock()
	atomic.StoreUint32(&o.done, 0)
	o.err = nil
}

// Error returns any error returned by the previous Do()
func (o *Once) Error() error {
	return o.err
}
