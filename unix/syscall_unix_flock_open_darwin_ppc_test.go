// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix_test

// flockOpenCloexec is the close-on-exec flag TestFcntlFlock passes to Open.
//
// Leopard's open(2) has no O_CLOEXEC (xnu-1228 bsd/sys/fcntl.h defines only
// FD_CLOEXEC, which fcntl(F_SETFD) sets after the open), so the package
// exports no such flag at this GOARCH and the test omits it. The test checks
// the fcntl(F_SETLK) calling convention, not close-on-exec.
const flockOpenCloexec = 0
