// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris) && !(darwin && (ppc || ppc64))

package unix_test

import "golang.org/x/sys/unix"

// flockOpenCloexec is the close-on-exec flag TestFcntlFlock passes to Open.
const flockOpenCloexec = unix.O_CLOEXEC
