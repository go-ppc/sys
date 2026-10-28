// Port-local (goppc-p9apt.22), not part of upstream golang.org/x/sys.
// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gc

#include "textflag.h"

// Syscall9 for darwin/ppc32. Package syscall has no Syscall9 here, so this is
// the trampoline itself, in the shape of syscall·Syscall6 in
// src/syscall/asm_darwin_ppc.s: the system call number is in R0, the first
// eight arguments are in R3-R10, `sc` resumes at the next instruction on
// failure (errno in R3) and skips it on success.
//
// The Leopard ppc kernel takes at most eight system call arguments in
// registers (xnu bsd/dev/ppc/munge.s has no munger wider than eight words and
// bsd/dev/ppc/systemcalls.c copies nothing from the user stack), and no
// system call it implements has more than eight. The ninth argument is
// therefore stored where the Darwin calling convention puts it, at
// FM_ARG0 = 56(R1) (xnu osfmk/ppc/asm.h), and is ignored by the kernel. It
// is delivered only to a callee that reads the stack, never to a system call.
// The frame is declared so that this slot belongs to Syscall9 rather than to
// its caller.

// func Syscall9(num, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err Errno)
TEXT	·Syscall9(SB),NOSPLIT,$64-52
	BL	runtime·entersyscall(SB)
	LWZ	R3, a9+36(FP)
	STW	R3, 56(R1)
	LWZ	R3, a1+4(FP)
	LWZ	R4, a2+8(FP)
	LWZ	R5, a3+12(FP)
	LWZ	R6, a4+16(FP)
	LWZ	R7, a5+20(FP)
	LWZ	R8, a6+24(FP)
	LWZ	R9, a7+28(FP)
	LWZ	R10, a8+32(FP)
	LWZ	R0, num+0(FP)
	SC
	B	error9
	XOR	R0, R0, R0
	STW	R3, r1+40(FP)
	STW	R4, r2+44(FP)
	STW	R0, err+48(FP)
	BL	runtime·exitsyscall(SB)
	RET
error9:
	XOR	R0, R0, R0
	STW	R3, err+48(FP)
	ADDI	R3, R0, $-1
	STW	R3, r1+40(FP)
	STW	R0, r2+44(FP)
	BL	runtime·exitsyscall(SB)
	RET
