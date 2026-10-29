// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gc

package unix_test

import (
	"os"
	"runtime"
	"syscall"
	"testing"
	"unsafe"

	"golang.org/x/sys/unix"
)

func TestLinuxPPCSyscall(t *testing.T) {
	pid, uid := os.Getpid(), os.Getuid()
	if got := unix.Getpid(); got != pid {
		t.Fatalf("Getpid = %d, want %d", got, pid)
	}
	if got := unix.Getuid(); got != uid {
		t.Fatalf("Getuid = %d, want %d", got, uid)
	}
	t.Logf("Getpid = %d, Getuid = %d (match os)", pid, uid)

	path := t.TempDir() + "/known"
	const contents = "linux/ppc syscall fixture\n"
	if err := os.WriteFile(path, []byte(contents), 0600); err != nil {
		t.Fatal(err)
	}
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	var st, fst unix.Stat_t
	if err := unix.Stat(path, &st); err != nil {
		t.Fatal(err)
	}
	if err := unix.Fstat(int(f.Fd()), &fst); err != nil {
		t.Fatal(err)
	}
	want, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}
	if st.Size != int64(len(contents)) || st.Size != want.Size() || st != fst {
		t.Fatalf("Stat = %+v, Fstat = %+v, os size = %d", st, fst, want.Size())
	}
	t.Logf("Stat/Fstat: size = %d, inode = %d, mode = %#o (match os size)", st.Size, st.Ino, st.Mode)

	missing := path + ".missing"
	if err := unix.Stat(missing, &st); err != unix.ENOENT {
		t.Errorf("Stat missing = %v, want ENOENT", err)
	}
	p, err := unix.BytePtrFromString(missing)
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		name string
		call func(uintptr, uintptr, uintptr, uintptr) (uintptr, uintptr, syscall.Errno)
	}{
		{"Syscall", unix.Syscall},
		{"RawSyscall", unix.RawSyscall},
		{"Syscall6", func(n, a1, a2, a3 uintptr) (uintptr, uintptr, syscall.Errno) {
			return unix.Syscall6(n, a1, a2, a3, 0, 0, 0)
		}},
		{"RawSyscall6", func(n, a1, a2, a3 uintptr) (uintptr, uintptr, syscall.Errno) {
			return unix.RawSyscall6(n, a1, a2, a3, 0, 0, 0)
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r1, r2, errno := tc.call(unix.SYS_GETPID, 0, 0, 0)
			t.Logf("GETPID: r1 = %d, r2 = %d, errno = %d", r1, r2, errno)
			if r1 != uintptr(pid) || errno != 0 {
				t.Errorf("GETPID = %d, errno %d; want %d, 0", r1, errno, pid)
			}
			r1, r2, errno = tc.call(unix.SYS_OPEN, uintptr(unsafe.Pointer(p)), unix.O_RDONLY, 0)
			runtime.KeepAlive(p)
			t.Logf("ENOENT: r1 = %#x, r2 = %#x, errno = %d", r1, r2, errno)
			if r1 != ^uintptr(0) || errno != unix.ENOENT {
				t.Errorf("OPEN missing = %#x, errno %d; want %#x, %d", r1, errno, ^uintptr(0), unix.ENOENT)
			}
		})
	}
	for _, tc := range []struct {
		name string
		call func(uintptr, uintptr, uintptr, uintptr) (uintptr, uintptr)
	}{
		{"SyscallNoError", unix.SyscallNoError},
		{"RawSyscallNoError", unix.RawSyscallNoError},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r1, _ := tc.call(unix.SYS_GETPID, 0, 0, 0)
			if r1 != uintptr(pid) {
				t.Errorf("GETPID = %d, want %d", r1, pid)
			}
			// Even failure must be returned raw by these entry points.
			r1, r2 := tc.call(unix.SYS_OPEN, uintptr(unsafe.Pointer(p)), unix.O_RDONLY, 0)
			runtime.KeepAlive(p)
			t.Logf("raw ENOENT: r1 = %d, r2 = %d", r1, r2)
			if r1 != uintptr(unix.ENOENT) || r2 != unix.O_RDONLY {
				t.Errorf("raw OPEN missing = (%d, %d), want (%d, %d)", r1, r2, unix.ENOENT, unix.O_RDONLY)
			}
		})
	}
}
