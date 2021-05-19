package resync_test

import (
	"testing"

	"github.com/szampardi/resync"
)

func TestOnceReset(t *testing.T) {
	var calls int
	var c = resync.Once{Fn: func() error {
		calls++
		return nil
	}}
	c.Do()
	c.Do()
	c.Do()
	t.Logf("calls should be 1 now: %d", calls)
	if calls != 1 {
		t.Fail()
	}
	c.Reset()
	c.Do()
	c.Do()
	c.Do()
	t.Logf("calls should be 2 now: %d", calls)
	if calls != 2 {
		t.Fail()
	}
}
