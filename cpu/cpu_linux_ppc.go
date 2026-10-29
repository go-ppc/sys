// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && ppc

package cpu

// cacheLineSize is a conservative padding hint, not a machine descriptor.
// Measured on powerbook-g4 (PowerBook5,6, Darwin 9.8.0): hw.cachelinesize=32,
// hw.l1icachesize=32768, hw.cpusubtype=11. Other ppc32 cores may differ.
// Use 128 to also cover the PowerPC 970's documented 128-byte cache line and
// match Go's internal/cpu ppc padding. Over-padding wastes space; under-padding
// can allow false sharing. See runs/xforks-linuxppc-20261029 for the measurement.
const cacheLineSize = 128

// No exported ppc32 feature flags exist for GODEBUG options to control.
func initOptions() {}
