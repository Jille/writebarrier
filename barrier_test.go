package writebarrier

import (
	"runtime"
	"testing"
)

func TestIsEnabled(t *testing.T) {
	finishedCycle := waitForGCEnd()
	var garbage []byte
	for !IsEnabled() {
		garbage = append(garbage, 0)
		escape = garbage
	}
	t.Logf("cycles: %d", len(garbage))
	<-finishedCycle
	if IsEnabled() {
		t.Error("unexpectedly still in GC")
	}
}

var escape any

func waitForGCEnd() chan struct{} {
	ch := make(chan struct{})
	sentinel := new([4096]byte)
	runtime.SetFinalizer(sentinel, func(s *[4096]byte) {
		close(ch)
	})
	escape = sentinel
	return ch
}
