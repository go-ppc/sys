// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin

package ifreqmtu_test

import "golang.org/x/sys/unix"

// Compile with go test -c -vet=off ./unix/testdata/ifreqmtu. This fixture is
// separate because vet normally discourages unkeyed literals of external types;
// here that syntax is precisely the source compatibility being tested.
var (
	name = [unix.IFNAMSIZ]byte{'l', 'o', '0'}
	mtu  = int32(1500)
	_    = unix.IfreqMTU{name, mtu}
)
