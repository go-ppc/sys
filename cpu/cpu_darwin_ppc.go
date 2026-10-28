// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin && ppc

package cpu

// cacheLineSize is the padding unit of CacheLinePad. It is a compile-time array
// length, so it cannot be read from the kernel at run time the way upstream's
// darwin files read feature flags with sysctlbyname (cpu_darwin_x86.go,
// cpu_darwin_arm64.go). The value is the larger of two sources:
//
//   - measured on powerbook-g4 (PowerBook5,6, Mac OS X 10.5.8, G4):
//     sysctlbyname("hw.cachelinesize") = 32 (rc=0, an 8-byte quad), as is
//     hw.l1dcachesize = 32768; run goppc-p9apt.24, runs/xsys-cpu-20261028.
//   - derived from reference/text/PPC970MP-manual-v2.3.txt:1893 (G5, "128-byte
//     cache line"). No G5 was executed; a 32-bit image also runs on the G5.
//
// 128 is also the value of std's internal/cpu/cpu_ppc.go, so both stay
// consistent. Over-padding costs memory only; 32 would under-pad on the G5.
const cacheLineSize = 128

// initOptions registers no GODEBUG options: this package has no exported
// feature struct for 32-bit PowerPC, so there is nothing for an option to set.
func initOptions() {}

// archInit detects nothing. AltiVec is visible to sysctlbyname as
// hw.optional.altivec (measured =1 on powerbook-g4, goppc-p9apt.24), but there
// is no exported field to hold it and adding one is an API change upstream has
// not made. Initialized stays false, as in cpu_other_arm.go.
func archInit() {}
