// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package unix_test

import (
	"os"
	"testing"

	"golang.org/x/sys/unix"
)

// skipIfPollCannotWatchFIFO skips the caller when this kernel refuses to
// watch a FIFO through poll(2) as a matter of policy, not of state.
//
// Darwin 9 is such a kernel. Its poll(2) is implemented over kqueue (xnu-1228
// bsd/kern/sys_generic.c poll_nocancel registers one kevent per pollfd), a
// FIFO vnode has no kqueue filter, and poll_nocancel turns the failed
// registration into an immediate revents = POLLNVAL for a FIFO fd that is open
// and valid. Darwin 10 added the FIFO branch to the vnode kqfilter, but GOOS
// stayed "darwin" across both, so the premise has to be probed at run time
// instead of being spelled as a build tag or a GOOS switch.
//
// The FIFO verdict alone does not justify the skip: a stale or wrong fd number
// also comes back POLLNVAL. The pipe(2) control fd polled in the same probe
// must come back clean, so a real regression of the wrapper fails here instead
// of quietly skipping the test.
func skipIfPollCannotWatchFIFO(t *testing.T, fifo *os.File) {
	t.Helper()

	fds := []unix.PollFd{{Fd: int32(fifo.Fd()), Events: unix.POLLIN}}
	if _, err := unix.Poll(fds, 0); err != nil {
		return // TestPoll reports the unexpected error in its own terms.
	}
	if fds[0].Revents&unix.POLLNVAL == 0 {
		return
	}

	var p [2]int
	if err := unix.Pipe(p[:]); err != nil {
		t.Fatalf("poll capability probe: Pipe: %v", err)
	}
	defer unix.Close(p[0])
	defer unix.Close(p[1])

	control := []unix.PollFd{{Fd: int32(p[0]), Events: unix.POLLIN}}
	if _, err := unix.Poll(control, 0); err != nil {
		t.Fatalf("poll capability probe: Poll of a pipe fd: %v", err)
	}
	if control[0].Revents&unix.POLLNVAL != 0 {
		t.Fatalf("poll capability probe: Poll reported POLLNVAL for a valid pipe fd; the FIFO POLLNVAL is therefore not this kernel's FIFO limitation")
	}

	t.Skip("poll(2) cannot watch a FIFO on this kernel: an open, valid FIFO fd is reported POLLNVAL immediately, as on Darwin 9 where poll is kqueue-backed and a FIFO vnode has no kqueue filter")
}
