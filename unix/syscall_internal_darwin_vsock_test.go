// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin && (amd64 || arm64 || ppc64)

package unix

import (
	"reflect"
	"testing"
	"unsafe"
)

// These tests use constants from the SDK-derived Darwin error tables on amd64
// and arm64, and the ppc64 source-compatibility declarations. Leopard's ppc SDK
// does not define the VMADDR constants.
func Test_anyToSockaddrVM_darwin(t *testing.T) {
	tests := []struct {
		name string
		rsa  *RawSockaddrAny
		sa   Sockaddr
		err  error
	}{
		{
			name: "AF_VSOCK empty",
			rsa:  sockaddrVMToAny(RawSockaddrVM{}),
			err:  EAFNOSUPPORT,
		},
		{
			name: "AF_VSOCK Cid and Port",
			rsa: sockaddrVMToAny(RawSockaddrVM{
				Family: AF_VSOCK,
				Cid:    VMADDR_CID_HOST,
				Port:   VMADDR_PORT_ANY,
			}),
			sa: &SockaddrVM{
				CID:  VMADDR_CID_HOST,
				Port: VMADDR_PORT_ANY,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fd := int(0)
			sa, err := anyToSockaddr(fd, tt.rsa)
			if err != tt.err {
				t.Fatalf("unexpected error: %v, want: %v", err, tt.err)
			}

			if !reflect.DeepEqual(sa, tt.sa) {
				t.Fatalf("unexpected Sockaddr:\n got: %#v\nwant: %#v", sa, tt.sa)
			}
		})
	}
}

func TestSockaddrVM_sockaddr(t *testing.T) {
	tests := []struct {
		name string
		sa   *SockaddrVM
		raw  *RawSockaddrVM
		err  error
	}{
		{
			name: "empty",
			sa:   &SockaddrVM{},
			raw: &RawSockaddrVM{
				Len:    SizeofSockaddrVM,
				Family: AF_VSOCK,
			},
		},
		{
			name: "with CID and port",
			sa: &SockaddrVM{
				CID:  VMADDR_CID_HOST,
				Port: VMADDR_PORT_ANY,
			},
			raw: &RawSockaddrVM{
				Len:    SizeofSockaddrVM,
				Family: AF_VSOCK,
				Port:   VMADDR_PORT_ANY,
				Cid:    VMADDR_CID_HOST,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, l, err := tt.sa.sockaddr()
			if err != tt.err {
				t.Fatalf("unexpected error: %v, want: %v", err, tt.err)
			}

			// Must be 0 on error or a fixed size otherwise.
			if (tt.err != nil && l != 0) || (tt.raw != nil && l != SizeofSockaddrVM) {
				t.Fatalf("unexpected Socklen: %d", l)
			}

			if out != nil {
				raw := (*RawSockaddrVM)(out)
				if !reflect.DeepEqual(raw, tt.raw) {
					t.Fatalf("unexpected RawSockaddrVM:\n got: %#v\nwant: %#v", raw, tt.raw)
				}
			}
		})
	}
}

func sockaddrVMToAny(in RawSockaddrVM) *RawSockaddrAny {
	var out RawSockaddrAny
	copy(
		(*(*[SizeofSockaddrAny]byte)(unsafe.Pointer(&out)))[:],
		(*(*[SizeofSockaddrVM]byte)(unsafe.Pointer(&in)))[:],
	)
	return &out
}
