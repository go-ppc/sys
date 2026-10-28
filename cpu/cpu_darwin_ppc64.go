// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin && ppc64

package cpu

// archInit replaces cpu_other_ppc64x.go on darwin. That stub sets
// PPC64.IsPOWER8 = true for every non-linux, non-aix OS, which is false on every
// Mac OS X ppc64 machine: the G5 (PowerPC 970) predates ISA v2.07. Nothing is
// detected here; Initialized = true keeps the stub's contract that the package
// is usable. cacheLineSize = 128 and initOptions come from cpu_ppc64x.go
// (transcribed from upstream; the 970MP manual, reference/text/
// PPC970MP-manual-v2.3.txt:1893, documents the same 128-byte line). Built
// only: no G5 was executed (pmg5 unreachable).
func archInit() {
	Initialized = true
}
