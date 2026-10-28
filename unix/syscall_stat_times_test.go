// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package unix_test

import "golang.org/x/sys/unix"

func statAtim(st *unix.Stat_t) *unix.Timespec { return &st.Atim }
func statMtim(st *unix.Stat_t) *unix.Timespec { return &st.Mtim }
func statCtim(st *unix.Stat_t) *unix.Timespec { return &st.Ctim }
