// Port-local (goppc-p9apt.22), not part of upstream golang.org/x/sys.
// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gc

#include "textflag.h"

//
// System call support for ppc64, Darwin
//

// asm_bsd_ppc64.s serves the other BSDs; Darwin has its own file because
// package syscall has no Syscall9 at darwin/ppc64, so the shared row
// `JMP syscall·Syscall9(SB)` would reference an undefined symbol.

// Just jump to package syscall's implementation for all these functions.
// The runtime may know about them.

TEXT	·Syscall(SB),NOSPLIT,$0-56
	JMP	syscall·Syscall(SB)

TEXT	·Syscall6(SB),NOSPLIT,$0-80
	JMP	syscall·Syscall6(SB)

TEXT	·RawSyscall(SB),NOSPLIT,$0-56
	JMP	syscall·RawSyscall(SB)

TEXT	·RawSyscall6(SB),NOSPLIT,$0-80
	JMP	syscall·RawSyscall6(SB)

// Syscall9 is the trampoline itself, in the shape of syscall·Syscall6 in
// src/syscall/asm_darwin_ppc64.s: the system call number is in R0, the first
// eight arguments are in R3-R10, `sc` resumes at the next instruction on
// failure (errno in R3) and two instructions on, with R3/R4 results, on
// success; SYSCALL is a bare sc so that those resume points do not move.
//
// The Leopard ppc kernel takes at most eight system call arguments in
// registers and copies nothing from the user stack (xnu bsd/dev/ppc/munge.s,
// bsd/dev/ppc/systemcalls.c), and no system call it implements has more than
// eight. The ninth argument is stored at 112(R1), the first parameter slot
// past the 48-byte linkage area and the 64-byte register parameter save area
// of the 64-bit Darwin PowerPC convention, and is ignored by the kernel. That
// offset follows the convention and has not been observed on a G5.

// func Syscall9(num, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err Errno)
TEXT	·Syscall9(SB),NOSPLIT,$128-104
	BL	runtime·entersyscall(SB)
	MOVD	a9+72(FP), R3
	MOVD	R3, 112(R1)
	MOVD	a1+8(FP), R3
	MOVD	a2+16(FP), R4
	MOVD	a3+24(FP), R5
	MOVD	a4+32(FP), R6
	MOVD	a5+40(FP), R7
	MOVD	a6+48(FP), R8
	MOVD	a7+56(FP), R9
	MOVD	a8+64(FP), R10
	MOVD	num+0(FP), R0
	SYSCALL
	BR	error9
	XOR	R0, R0, R0
	MOVD	R3, r1+80(FP)
	MOVD	R4, r2+88(FP)
	MOVD	R0, err+96(FP)
	BL	runtime·exitsyscall(SB)
	RET
error9:
	XOR	R0, R0, R0
	MOVD	R3, err+96(FP)
	MOVD	$-1, R3
	MOVD	R3, r1+80(FP)
	MOVD	R0, r2+88(FP)
	BL	runtime·exitsyscall(SB)
	RET
