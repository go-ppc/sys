// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin

package unix

import "unsafe"

// struct ifreq on Darwin, read out of the 10.5.8 headers by Apple gcc 4.0.1
// build 5493: IFNAMSIZ bytes of ifr_name, then a 16-byte union whose first int
// is ifr_mtu. Measured with -m32 and asserted with -m64 on a PowerBook5,6; the
// measurement and its -m64 and mutant controls are in
// runs/xsys-ifreqmtu-20261029/.
const (
	darwinIfreqSize       = 32
	darwinIfreqNameOffset = 0
	darwinIfreqMTUOffset  = 16
)

// The public type retains upstream's size and offsets; only the private ioctl
// argument spans the full union. Each equality is checked in both directions.
// No ioctl is issued here.
var (
	_ [20 - int(unsafe.Sizeof(IfreqMTU{}))]byte
	_ [int(unsafe.Sizeof(IfreqMTU{})) - 20]byte

	_ [darwinIfreqNameOffset - int(unsafe.Offsetof(IfreqMTU{}.Name))]byte
	_ [int(unsafe.Offsetof(IfreqMTU{}.Name)) - darwinIfreqNameOffset]byte

	_ [darwinIfreqMTUOffset - int(unsafe.Offsetof(IfreqMTU{}.MTU))]byte
	_ [int(unsafe.Offsetof(IfreqMTU{}.MTU)) - darwinIfreqMTUOffset]byte

	_ [darwinIfreqSize - int(unsafe.Sizeof(ifreqMTUArgs{}))]byte
	_ [int(unsafe.Sizeof(ifreqMTUArgs{})) - darwinIfreqSize]byte

	_ [darwinIfreqNameOffset - int(unsafe.Offsetof(ifreqMTUArgs{}.Name))]byte
	_ [int(unsafe.Offsetof(ifreqMTUArgs{}.Name)) - darwinIfreqNameOffset]byte

	_ [darwinIfreqMTUOffset - int(unsafe.Offsetof(ifreqMTUArgs{}.MTU))]byte
	_ [int(unsafe.Offsetof(ifreqMTUArgs{}.MTU)) - darwinIfreqMTUOffset]byte
)
