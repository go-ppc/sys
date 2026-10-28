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

// Each property is checked in both directions: a negative array length does not
// compile, a zero-length one is fine, so a layout that is wrong on only one
// side still fails. SIOCGIFMTU's own encoded length is tied to the measured 32
// as well, so a change in the generated tables cannot satisfy this silently.
//
// No ioctl is issued here: that would copy the full 32 bytes through the type,
// and SIOCSIFMTU would change a live interface's MTU.
var (
	_ [darwinIfreqSize - int(unsafe.Sizeof(IfreqMTU{}))]byte
	_ [int(unsafe.Sizeof(IfreqMTU{})) - darwinIfreqSize]byte

	_ [darwinIfreqNameOffset - int(unsafe.Offsetof(IfreqMTU{}.Name))]byte
	_ [int(unsafe.Offsetof(IfreqMTU{}.Name)) - darwinIfreqNameOffset]byte

	_ [darwinIfreqMTUOffset - int(unsafe.Offsetof(IfreqMTU{}.MTU))]byte
	_ [int(unsafe.Offsetof(IfreqMTU{}.MTU)) - darwinIfreqMTUOffset]byte

	_ [darwinIfreqSize - ifreqMTUSize]byte
	_ [ifreqMTUSize - darwinIfreqSize]byte
)
