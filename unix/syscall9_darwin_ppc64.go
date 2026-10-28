// Port-local (goppc-p9apt.22), not part of upstream golang.org/x/sys.
// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin && ppc64 && gc

package unix

import "syscall"

// Syscall9 is implemented in asm_darwin_ppc64.s. Package syscall has no
// Syscall9 at darwin/ppc64, so there is nothing to jump to.
func Syscall9(num, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err syscall.Errno)
