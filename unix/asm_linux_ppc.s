// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build gc

#include "textflag.h"

// Linux/ppc passes the syscall number in R0, arguments in R3-R8, and
// results in R3-R4. After SC, CR0.SO set means failure with a positive
// errno in R3 (not -errno). Translate that to r1 = -1 and err = errno.
// The NoError entry points deliberately return R3-R4 without translation.

TEXT ·Syscall(SB),NOSPLIT,$0-28
	BL runtime·entersyscall(SB)
	LWZ R3, a1+4(FP)
	LWZ R4, a2+8(FP)
	LWZ R5, a3+12(FP)
	ADDI R6, R0, $0
	ADDI R7, R0, $0
	ADDI R8, R0, $0
	LWZ R0, trap+0(FP)
	SC
	BC 4, 3, syscall_ok
	STW R3, err+24(FP)
	ADDI R3, R0, $-1
	B syscall_done
syscall_ok:
	ADDI R0, R0, $0
	STW R0, err+24(FP)
syscall_done:
	STW R3, r1+16(FP)
	STW R4, r2+20(FP)
	BL runtime·exitsyscall(SB)
	RET

TEXT ·Syscall6(SB),NOSPLIT,$0-40
	BL runtime·entersyscall(SB)
	LWZ R3, a1+4(FP)
	LWZ R4, a2+8(FP)
	LWZ R5, a3+12(FP)
	LWZ R6, a4+16(FP)
	LWZ R7, a5+20(FP)
	LWZ R8, a6+24(FP)
	LWZ R0, trap+0(FP)
	SC
	BC 4, 3, syscall6_ok
	STW R3, err+36(FP)
	ADDI R3, R0, $-1
	B syscall6_done
syscall6_ok:
	ADDI R0, R0, $0
	STW R0, err+36(FP)
syscall6_done:
	STW R3, r1+28(FP)
	STW R4, r2+32(FP)
	BL runtime·exitsyscall(SB)
	RET

TEXT ·RawSyscall(SB),NOSPLIT|NOFRAME,$0-28
	LWZ R3, a1+4(FP)
	LWZ R4, a2+8(FP)
	LWZ R5, a3+12(FP)
	ADDI R6, R0, $0
	ADDI R7, R0, $0
	ADDI R8, R0, $0
	LWZ R0, trap+0(FP)
	SC
	BC 4, 3, rawsyscall_ok
	STW R3, err+24(FP)
	ADDI R3, R0, $-1
	B rawsyscall_done
rawsyscall_ok:
	ADDI R0, R0, $0
	STW R0, err+24(FP)
rawsyscall_done:
	STW R3, r1+16(FP)
	STW R4, r2+20(FP)
	RET

TEXT ·RawSyscall6(SB),NOSPLIT|NOFRAME,$0-40
	LWZ R3, a1+4(FP)
	LWZ R4, a2+8(FP)
	LWZ R5, a3+12(FP)
	LWZ R6, a4+16(FP)
	LWZ R7, a5+20(FP)
	LWZ R8, a6+24(FP)
	LWZ R0, trap+0(FP)
	SC
	BC 4, 3, rawsyscall6_ok
	STW R3, err+36(FP)
	ADDI R3, R0, $-1
	B rawsyscall6_done
rawsyscall6_ok:
	ADDI R0, R0, $0
	STW R0, err+36(FP)
rawsyscall6_done:
	STW R3, r1+28(FP)
	STW R4, r2+32(FP)
	RET

TEXT ·SyscallNoError(SB),NOSPLIT,$0-24
	BL runtime·entersyscall(SB)
	LWZ R3, a1+4(FP)
	LWZ R4, a2+8(FP)
	LWZ R5, a3+12(FP)
	ADDI R6, R0, $0
	ADDI R7, R0, $0
	ADDI R8, R0, $0
	LWZ R0, trap+0(FP)
	SC
	STW R3, r1+16(FP)
	STW R4, r2+20(FP)
	BL runtime·exitsyscall(SB)
	RET

TEXT ·RawSyscallNoError(SB),NOSPLIT|NOFRAME,$0-24
	LWZ R3, a1+4(FP)
	LWZ R4, a2+8(FP)
	LWZ R5, a3+12(FP)
	ADDI R6, R0, $0
	ADDI R7, R0, $0
	ADDI R8, R0, $0
	LWZ R0, trap+0(FP)
	SC
	STW R3, r1+16(FP)
	STW R4, r2+20(FP)
	RET
