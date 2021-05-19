package resync_test

import (
	"fmt"
	"testing"
)

func TestOnceReset(t *testing.T) {
	var calls int
	var c = Once{Fn: func() error {
		calls++
		return nil
	}}
	c.Do()
	c.Do()
	c.Do()
	fmt.Println("calls should be 1 now:", calls)
	if calls != 1 {
		t.Fail()
	}
	c.Reset()
	c.Do()
	c.Do()
	c.Do()
	fmt.Println("calls should be 2 now:", calls)
	if calls != 2 {
		t.Fail()
	}
}
