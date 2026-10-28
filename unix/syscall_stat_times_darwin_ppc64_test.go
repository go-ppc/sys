// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin

package unix_test

import "golang.org/x/sys/unix"

// Darwin's ppc64 Stat_t follows the C st_atimespec/st_mtimespec/st_ctimespec names.
func statAtim(st *unix.Stat_t) *unix.Timespec { return &st.Atimespec }
func statMtim(st *unix.Stat_t) *unix.Timespec { return &st.Mtimespec }
func statCtim(st *unix.Stat_t) *unix.Timespec { return &st.Ctimespec }
