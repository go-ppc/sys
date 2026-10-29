// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import "unsafe"

// Route layouts from Apple gcc 4.0.1 build 5493 on Leopard 10.5.8:
// C ppc32 executed, C ppc64 compile-only, Go both widths built and inspected.
// Evidence and failing controls: runs/xsys-routelayout-20261029/.
// Every size, named field offset, alignment and leading-byte embedding is
// checked in both directions. Nested timeval offsets are checked as well.
// IfData64 and IfMsghdr2 deliberately assert Go alignment/embedding 8 here,
// NOT C's 4: net/if_var.h packs if_data64 to 4. Flat sizes and offsets match;
// these guards document the remaining alignment limitation, not a fix for it.
var (
	// IfMsghdr.size
	_ [112 - int(unsafe.Sizeof(IfMsghdr{}))]byte
	_ [int(unsafe.Sizeof(IfMsghdr{})) - 112]byte
	// IfMsghdr.align
	_ [4 - int(unsafe.Alignof(IfMsghdr{}))]byte
	_ [int(unsafe.Alignof(IfMsghdr{})) - 4]byte
	// IfMsghdr.Msglen
	_ [0 - int(unsafe.Offsetof(IfMsghdr{}.Msglen))]byte
	_ [int(unsafe.Offsetof(IfMsghdr{}.Msglen)) - 0]byte
	// IfMsghdr.Version
	_ [2 - int(unsafe.Offsetof(IfMsghdr{}.Version))]byte
	_ [int(unsafe.Offsetof(IfMsghdr{}.Version)) - 2]byte
	// IfMsghdr.Type
	_ [3 - int(unsafe.Offsetof(IfMsghdr{}.Type))]byte
	_ [int(unsafe.Offsetof(IfMsghdr{}.Type)) - 3]byte
	// IfMsghdr.Addrs
	_ [4 - int(unsafe.Offsetof(IfMsghdr{}.Addrs))]byte
	_ [int(unsafe.Offsetof(IfMsghdr{}.Addrs)) - 4]byte
	// IfMsghdr.Flags
	_ [8 - int(unsafe.Offsetof(IfMsghdr{}.Flags))]byte
	_ [int(unsafe.Offsetof(IfMsghdr{}.Flags)) - 8]byte
	// IfMsghdr.Index
	_ [12 - int(unsafe.Offsetof(IfMsghdr{}.Index))]byte
	_ [int(unsafe.Offsetof(IfMsghdr{}.Index)) - 12]byte
	// IfMsghdr.Data
	_ [16 - int(unsafe.Offsetof(IfMsghdr{}.Data))]byte
	_ [int(unsafe.Offsetof(IfMsghdr{}.Data)) - 16]byte
	// IfMsghdr2.size
	_ [160 - int(unsafe.Sizeof(IfMsghdr2{}))]byte
	_ [int(unsafe.Sizeof(IfMsghdr2{})) - 160]byte
	// IfMsghdr2.align
	_ [8 - int(unsafe.Alignof(IfMsghdr2{}))]byte
	_ [int(unsafe.Alignof(IfMsghdr2{})) - 8]byte
	// IfMsghdr2.Msglen
	_ [0 - int(unsafe.Offsetof(IfMsghdr2{}.Msglen))]byte
	_ [int(unsafe.Offsetof(IfMsghdr2{}.Msglen)) - 0]byte
	// IfMsghdr2.Version
	_ [2 - int(unsafe.Offsetof(IfMsghdr2{}.Version))]byte
	_ [int(unsafe.Offsetof(IfMsghdr2{}.Version)) - 2]byte
	// IfMsghdr2.Type
	_ [3 - int(unsafe.Offsetof(IfMsghdr2{}.Type))]byte
	_ [int(unsafe.Offsetof(IfMsghdr2{}.Type)) - 3]byte
	// IfMsghdr2.Addrs
	_ [4 - int(unsafe.Offsetof(IfMsghdr2{}.Addrs))]byte
	_ [int(unsafe.Offsetof(IfMsghdr2{}.Addrs)) - 4]byte
	// IfMsghdr2.Flags
	_ [8 - int(unsafe.Offsetof(IfMsghdr2{}.Flags))]byte
	_ [int(unsafe.Offsetof(IfMsghdr2{}.Flags)) - 8]byte
	// IfMsghdr2.Index
	_ [12 - int(unsafe.Offsetof(IfMsghdr2{}.Index))]byte
	_ [int(unsafe.Offsetof(IfMsghdr2{}.Index)) - 12]byte
	// IfMsghdr2.Snd_len
	_ [16 - int(unsafe.Offsetof(IfMsghdr2{}.Snd_len))]byte
	_ [int(unsafe.Offsetof(IfMsghdr2{}.Snd_len)) - 16]byte
	// IfMsghdr2.Snd_maxlen
	_ [20 - int(unsafe.Offsetof(IfMsghdr2{}.Snd_maxlen))]byte
	_ [int(unsafe.Offsetof(IfMsghdr2{}.Snd_maxlen)) - 20]byte
	// IfMsghdr2.Snd_drops
	_ [24 - int(unsafe.Offsetof(IfMsghdr2{}.Snd_drops))]byte
	_ [int(unsafe.Offsetof(IfMsghdr2{}.Snd_drops)) - 24]byte
	// IfMsghdr2.Timer
	_ [28 - int(unsafe.Offsetof(IfMsghdr2{}.Timer))]byte
	_ [int(unsafe.Offsetof(IfMsghdr2{}.Timer)) - 28]byte
	// IfMsghdr2.Data
	_ [32 - int(unsafe.Offsetof(IfMsghdr2{}.Data))]byte
	_ [int(unsafe.Offsetof(IfMsghdr2{}.Data)) - 32]byte
	// IfData.size
	_ [96 - int(unsafe.Sizeof(IfData{}))]byte
	_ [int(unsafe.Sizeof(IfData{})) - 96]byte
	// IfData.align
	_ [4 - int(unsafe.Alignof(IfData{}))]byte
	_ [int(unsafe.Alignof(IfData{})) - 4]byte
	// IfData.Type
	_ [0 - int(unsafe.Offsetof(IfData{}.Type))]byte
	_ [int(unsafe.Offsetof(IfData{}.Type)) - 0]byte
	// IfData.Typelen
	_ [1 - int(unsafe.Offsetof(IfData{}.Typelen))]byte
	_ [int(unsafe.Offsetof(IfData{}.Typelen)) - 1]byte
	// IfData.Physical
	_ [2 - int(unsafe.Offsetof(IfData{}.Physical))]byte
	_ [int(unsafe.Offsetof(IfData{}.Physical)) - 2]byte
	// IfData.Addrlen
	_ [3 - int(unsafe.Offsetof(IfData{}.Addrlen))]byte
	_ [int(unsafe.Offsetof(IfData{}.Addrlen)) - 3]byte
	// IfData.Hdrlen
	_ [4 - int(unsafe.Offsetof(IfData{}.Hdrlen))]byte
	_ [int(unsafe.Offsetof(IfData{}.Hdrlen)) - 4]byte
	// IfData.Recvquota
	_ [5 - int(unsafe.Offsetof(IfData{}.Recvquota))]byte
	_ [int(unsafe.Offsetof(IfData{}.Recvquota)) - 5]byte
	// IfData.Xmitquota
	_ [6 - int(unsafe.Offsetof(IfData{}.Xmitquota))]byte
	_ [int(unsafe.Offsetof(IfData{}.Xmitquota)) - 6]byte
	// IfData.Unused1
	_ [7 - int(unsafe.Offsetof(IfData{}.Unused1))]byte
	_ [int(unsafe.Offsetof(IfData{}.Unused1)) - 7]byte
	// IfData.Mtu
	_ [8 - int(unsafe.Offsetof(IfData{}.Mtu))]byte
	_ [int(unsafe.Offsetof(IfData{}.Mtu)) - 8]byte
	// IfData.Metric
	_ [12 - int(unsafe.Offsetof(IfData{}.Metric))]byte
	_ [int(unsafe.Offsetof(IfData{}.Metric)) - 12]byte
	// IfData.Baudrate
	_ [16 - int(unsafe.Offsetof(IfData{}.Baudrate))]byte
	_ [int(unsafe.Offsetof(IfData{}.Baudrate)) - 16]byte
	// IfData.Ipackets
	_ [20 - int(unsafe.Offsetof(IfData{}.Ipackets))]byte
	_ [int(unsafe.Offsetof(IfData{}.Ipackets)) - 20]byte
	// IfData.Ierrors
	_ [24 - int(unsafe.Offsetof(IfData{}.Ierrors))]byte
	_ [int(unsafe.Offsetof(IfData{}.Ierrors)) - 24]byte
	// IfData.Opackets
	_ [28 - int(unsafe.Offsetof(IfData{}.Opackets))]byte
	_ [int(unsafe.Offsetof(IfData{}.Opackets)) - 28]byte
	// IfData.Oerrors
	_ [32 - int(unsafe.Offsetof(IfData{}.Oerrors))]byte
	_ [int(unsafe.Offsetof(IfData{}.Oerrors)) - 32]byte
	// IfData.Collisions
	_ [36 - int(unsafe.Offsetof(IfData{}.Collisions))]byte
	_ [int(unsafe.Offsetof(IfData{}.Collisions)) - 36]byte
	// IfData.Ibytes
	_ [40 - int(unsafe.Offsetof(IfData{}.Ibytes))]byte
	_ [int(unsafe.Offsetof(IfData{}.Ibytes)) - 40]byte
	// IfData.Obytes
	_ [44 - int(unsafe.Offsetof(IfData{}.Obytes))]byte
	_ [int(unsafe.Offsetof(IfData{}.Obytes)) - 44]byte
	// IfData.Imcasts
	_ [48 - int(unsafe.Offsetof(IfData{}.Imcasts))]byte
	_ [int(unsafe.Offsetof(IfData{}.Imcasts)) - 48]byte
	// IfData.Omcasts
	_ [52 - int(unsafe.Offsetof(IfData{}.Omcasts))]byte
	_ [int(unsafe.Offsetof(IfData{}.Omcasts)) - 52]byte
	// IfData.Iqdrops
	_ [56 - int(unsafe.Offsetof(IfData{}.Iqdrops))]byte
	_ [int(unsafe.Offsetof(IfData{}.Iqdrops)) - 56]byte
	// IfData.Noproto
	_ [60 - int(unsafe.Offsetof(IfData{}.Noproto))]byte
	_ [int(unsafe.Offsetof(IfData{}.Noproto)) - 60]byte
	// IfData.Recvtiming
	_ [64 - int(unsafe.Offsetof(IfData{}.Recvtiming))]byte
	_ [int(unsafe.Offsetof(IfData{}.Recvtiming)) - 64]byte
	// IfData.Xmittiming
	_ [68 - int(unsafe.Offsetof(IfData{}.Xmittiming))]byte
	_ [int(unsafe.Offsetof(IfData{}.Xmittiming)) - 68]byte
	// IfData.Lastchange
	_ [72 - int(unsafe.Offsetof(IfData{}.Lastchange))]byte
	_ [int(unsafe.Offsetof(IfData{}.Lastchange)) - 72]byte
	// IfData.Unused2
	_ [80 - int(unsafe.Offsetof(IfData{}.Unused2))]byte
	_ [int(unsafe.Offsetof(IfData{}.Unused2)) - 80]byte
	// IfData.Hwassist
	_ [84 - int(unsafe.Offsetof(IfData{}.Hwassist))]byte
	_ [int(unsafe.Offsetof(IfData{}.Hwassist)) - 84]byte
	// IfData.Reserved1
	_ [88 - int(unsafe.Offsetof(IfData{}.Reserved1))]byte
	_ [int(unsafe.Offsetof(IfData{}.Reserved1)) - 88]byte
	// IfData.Reserved2
	_ [92 - int(unsafe.Offsetof(IfData{}.Reserved2))]byte
	_ [int(unsafe.Offsetof(IfData{}.Reserved2)) - 92]byte
	// IfData.Lastchange.Sec
	_ [72 - int(unsafe.Offsetof(IfData{}.Lastchange)+unsafe.Offsetof(IfData{}.Lastchange.Sec))]byte
	_ [int(unsafe.Offsetof(IfData{}.Lastchange)+unsafe.Offsetof(IfData{}.Lastchange.Sec)) - 72]byte
	// IfData.Lastchange.Usec
	_ [76 - int(unsafe.Offsetof(IfData{}.Lastchange)+unsafe.Offsetof(IfData{}.Lastchange.Usec))]byte
	_ [int(unsafe.Offsetof(IfData{}.Lastchange)+unsafe.Offsetof(IfData{}.Lastchange.Usec)) - 76]byte
	// IfData64.size
	_ [128 - int(unsafe.Sizeof(IfData64{}))]byte
	_ [int(unsafe.Sizeof(IfData64{})) - 128]byte
	// IfData64.align
	_ [8 - int(unsafe.Alignof(IfData64{}))]byte
	_ [int(unsafe.Alignof(IfData64{})) - 8]byte
	// IfData64.Type
	_ [0 - int(unsafe.Offsetof(IfData64{}.Type))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Type)) - 0]byte
	// IfData64.Typelen
	_ [1 - int(unsafe.Offsetof(IfData64{}.Typelen))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Typelen)) - 1]byte
	// IfData64.Physical
	_ [2 - int(unsafe.Offsetof(IfData64{}.Physical))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Physical)) - 2]byte
	// IfData64.Addrlen
	_ [3 - int(unsafe.Offsetof(IfData64{}.Addrlen))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Addrlen)) - 3]byte
	// IfData64.Hdrlen
	_ [4 - int(unsafe.Offsetof(IfData64{}.Hdrlen))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Hdrlen)) - 4]byte
	// IfData64.Recvquota
	_ [5 - int(unsafe.Offsetof(IfData64{}.Recvquota))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Recvquota)) - 5]byte
	// IfData64.Xmitquota
	_ [6 - int(unsafe.Offsetof(IfData64{}.Xmitquota))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Xmitquota)) - 6]byte
	// IfData64.Unused1
	_ [7 - int(unsafe.Offsetof(IfData64{}.Unused1))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Unused1)) - 7]byte
	// IfData64.Mtu
	_ [8 - int(unsafe.Offsetof(IfData64{}.Mtu))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Mtu)) - 8]byte
	// IfData64.Metric
	_ [12 - int(unsafe.Offsetof(IfData64{}.Metric))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Metric)) - 12]byte
	// IfData64.Baudrate
	_ [16 - int(unsafe.Offsetof(IfData64{}.Baudrate))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Baudrate)) - 16]byte
	// IfData64.Ipackets
	_ [24 - int(unsafe.Offsetof(IfData64{}.Ipackets))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Ipackets)) - 24]byte
	// IfData64.Ierrors
	_ [32 - int(unsafe.Offsetof(IfData64{}.Ierrors))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Ierrors)) - 32]byte
	// IfData64.Opackets
	_ [40 - int(unsafe.Offsetof(IfData64{}.Opackets))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Opackets)) - 40]byte
	// IfData64.Oerrors
	_ [48 - int(unsafe.Offsetof(IfData64{}.Oerrors))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Oerrors)) - 48]byte
	// IfData64.Collisions
	_ [56 - int(unsafe.Offsetof(IfData64{}.Collisions))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Collisions)) - 56]byte
	// IfData64.Ibytes
	_ [64 - int(unsafe.Offsetof(IfData64{}.Ibytes))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Ibytes)) - 64]byte
	// IfData64.Obytes
	_ [72 - int(unsafe.Offsetof(IfData64{}.Obytes))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Obytes)) - 72]byte
	// IfData64.Imcasts
	_ [80 - int(unsafe.Offsetof(IfData64{}.Imcasts))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Imcasts)) - 80]byte
	// IfData64.Omcasts
	_ [88 - int(unsafe.Offsetof(IfData64{}.Omcasts))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Omcasts)) - 88]byte
	// IfData64.Iqdrops
	_ [96 - int(unsafe.Offsetof(IfData64{}.Iqdrops))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Iqdrops)) - 96]byte
	// IfData64.Noproto
	_ [104 - int(unsafe.Offsetof(IfData64{}.Noproto))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Noproto)) - 104]byte
	// IfData64.Recvtiming
	_ [112 - int(unsafe.Offsetof(IfData64{}.Recvtiming))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Recvtiming)) - 112]byte
	// IfData64.Xmittiming
	_ [116 - int(unsafe.Offsetof(IfData64{}.Xmittiming))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Xmittiming)) - 116]byte
	// IfData64.Lastchange
	_ [120 - int(unsafe.Offsetof(IfData64{}.Lastchange))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Lastchange)) - 120]byte
	// IfData64.Lastchange.Sec
	_ [120 - int(unsafe.Offsetof(IfData64{}.Lastchange)+unsafe.Offsetof(IfData64{}.Lastchange.Sec))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Lastchange)+unsafe.Offsetof(IfData64{}.Lastchange.Sec)) - 120]byte
	// IfData64.Lastchange.Usec
	_ [124 - int(unsafe.Offsetof(IfData64{}.Lastchange)+unsafe.Offsetof(IfData64{}.Lastchange.Usec))]byte
	_ [int(unsafe.Offsetof(IfData64{}.Lastchange)+unsafe.Offsetof(IfData64{}.Lastchange.Usec)) - 124]byte
	// IfaMsghdr.size
	_ [20 - int(unsafe.Sizeof(IfaMsghdr{}))]byte
	_ [int(unsafe.Sizeof(IfaMsghdr{})) - 20]byte
	// IfaMsghdr.align
	_ [4 - int(unsafe.Alignof(IfaMsghdr{}))]byte
	_ [int(unsafe.Alignof(IfaMsghdr{})) - 4]byte
	// IfaMsghdr.Msglen
	_ [0 - int(unsafe.Offsetof(IfaMsghdr{}.Msglen))]byte
	_ [int(unsafe.Offsetof(IfaMsghdr{}.Msglen)) - 0]byte
	// IfaMsghdr.Version
	_ [2 - int(unsafe.Offsetof(IfaMsghdr{}.Version))]byte
	_ [int(unsafe.Offsetof(IfaMsghdr{}.Version)) - 2]byte
	// IfaMsghdr.Type
	_ [3 - int(unsafe.Offsetof(IfaMsghdr{}.Type))]byte
	_ [int(unsafe.Offsetof(IfaMsghdr{}.Type)) - 3]byte
	// IfaMsghdr.Addrs
	_ [4 - int(unsafe.Offsetof(IfaMsghdr{}.Addrs))]byte
	_ [int(unsafe.Offsetof(IfaMsghdr{}.Addrs)) - 4]byte
	// IfaMsghdr.Flags
	_ [8 - int(unsafe.Offsetof(IfaMsghdr{}.Flags))]byte
	_ [int(unsafe.Offsetof(IfaMsghdr{}.Flags)) - 8]byte
	// IfaMsghdr.Index
	_ [12 - int(unsafe.Offsetof(IfaMsghdr{}.Index))]byte
	_ [int(unsafe.Offsetof(IfaMsghdr{}.Index)) - 12]byte
	// IfaMsghdr.Metric
	_ [16 - int(unsafe.Offsetof(IfaMsghdr{}.Metric))]byte
	_ [int(unsafe.Offsetof(IfaMsghdr{}.Metric)) - 16]byte
	// IfmaMsghdr.size
	_ [16 - int(unsafe.Sizeof(IfmaMsghdr{}))]byte
	_ [int(unsafe.Sizeof(IfmaMsghdr{})) - 16]byte
	// IfmaMsghdr.align
	_ [4 - int(unsafe.Alignof(IfmaMsghdr{}))]byte
	_ [int(unsafe.Alignof(IfmaMsghdr{})) - 4]byte
	// IfmaMsghdr.Msglen
	_ [0 - int(unsafe.Offsetof(IfmaMsghdr{}.Msglen))]byte
	_ [int(unsafe.Offsetof(IfmaMsghdr{}.Msglen)) - 0]byte
	// IfmaMsghdr.Version
	_ [2 - int(unsafe.Offsetof(IfmaMsghdr{}.Version))]byte
	_ [int(unsafe.Offsetof(IfmaMsghdr{}.Version)) - 2]byte
	// IfmaMsghdr.Type
	_ [3 - int(unsafe.Offsetof(IfmaMsghdr{}.Type))]byte
	_ [int(unsafe.Offsetof(IfmaMsghdr{}.Type)) - 3]byte
	// IfmaMsghdr.Addrs
	_ [4 - int(unsafe.Offsetof(IfmaMsghdr{}.Addrs))]byte
	_ [int(unsafe.Offsetof(IfmaMsghdr{}.Addrs)) - 4]byte
	// IfmaMsghdr.Flags
	_ [8 - int(unsafe.Offsetof(IfmaMsghdr{}.Flags))]byte
	_ [int(unsafe.Offsetof(IfmaMsghdr{}.Flags)) - 8]byte
	// IfmaMsghdr.Index
	_ [12 - int(unsafe.Offsetof(IfmaMsghdr{}.Index))]byte
	_ [int(unsafe.Offsetof(IfmaMsghdr{}.Index)) - 12]byte
	// IfmaMsghdr2.size
	_ [20 - int(unsafe.Sizeof(IfmaMsghdr2{}))]byte
	_ [int(unsafe.Sizeof(IfmaMsghdr2{})) - 20]byte
	// IfmaMsghdr2.align
	_ [4 - int(unsafe.Alignof(IfmaMsghdr2{}))]byte
	_ [int(unsafe.Alignof(IfmaMsghdr2{})) - 4]byte
	// IfmaMsghdr2.Msglen
	_ [0 - int(unsafe.Offsetof(IfmaMsghdr2{}.Msglen))]byte
	_ [int(unsafe.Offsetof(IfmaMsghdr2{}.Msglen)) - 0]byte
	// IfmaMsghdr2.Version
	_ [2 - int(unsafe.Offsetof(IfmaMsghdr2{}.Version))]byte
	_ [int(unsafe.Offsetof(IfmaMsghdr2{}.Version)) - 2]byte
	// IfmaMsghdr2.Type
	_ [3 - int(unsafe.Offsetof(IfmaMsghdr2{}.Type))]byte
	_ [int(unsafe.Offsetof(IfmaMsghdr2{}.Type)) - 3]byte
	// IfmaMsghdr2.Addrs
	_ [4 - int(unsafe.Offsetof(IfmaMsghdr2{}.Addrs))]byte
	_ [int(unsafe.Offsetof(IfmaMsghdr2{}.Addrs)) - 4]byte
	// IfmaMsghdr2.Flags
	_ [8 - int(unsafe.Offsetof(IfmaMsghdr2{}.Flags))]byte
	_ [int(unsafe.Offsetof(IfmaMsghdr2{}.Flags)) - 8]byte
	// IfmaMsghdr2.Index
	_ [12 - int(unsafe.Offsetof(IfmaMsghdr2{}.Index))]byte
	_ [int(unsafe.Offsetof(IfmaMsghdr2{}.Index)) - 12]byte
	// IfmaMsghdr2.Refcount
	_ [16 - int(unsafe.Offsetof(IfmaMsghdr2{}.Refcount))]byte
	_ [int(unsafe.Offsetof(IfmaMsghdr2{}.Refcount)) - 16]byte
	// RtMsghdr.size
	_ [92 - int(unsafe.Sizeof(RtMsghdr{}))]byte
	_ [int(unsafe.Sizeof(RtMsghdr{})) - 92]byte
	// RtMsghdr.align
	_ [4 - int(unsafe.Alignof(RtMsghdr{}))]byte
	_ [int(unsafe.Alignof(RtMsghdr{})) - 4]byte
	// RtMsghdr.Msglen
	_ [0 - int(unsafe.Offsetof(RtMsghdr{}.Msglen))]byte
	_ [int(unsafe.Offsetof(RtMsghdr{}.Msglen)) - 0]byte
	// RtMsghdr.Version
	_ [2 - int(unsafe.Offsetof(RtMsghdr{}.Version))]byte
	_ [int(unsafe.Offsetof(RtMsghdr{}.Version)) - 2]byte
	// RtMsghdr.Type
	_ [3 - int(unsafe.Offsetof(RtMsghdr{}.Type))]byte
	_ [int(unsafe.Offsetof(RtMsghdr{}.Type)) - 3]byte
	// RtMsghdr.Index
	_ [4 - int(unsafe.Offsetof(RtMsghdr{}.Index))]byte
	_ [int(unsafe.Offsetof(RtMsghdr{}.Index)) - 4]byte
	// RtMsghdr.Flags
	_ [8 - int(unsafe.Offsetof(RtMsghdr{}.Flags))]byte
	_ [int(unsafe.Offsetof(RtMsghdr{}.Flags)) - 8]byte
	// RtMsghdr.Addrs
	_ [12 - int(unsafe.Offsetof(RtMsghdr{}.Addrs))]byte
	_ [int(unsafe.Offsetof(RtMsghdr{}.Addrs)) - 12]byte
	// RtMsghdr.Pid
	_ [16 - int(unsafe.Offsetof(RtMsghdr{}.Pid))]byte
	_ [int(unsafe.Offsetof(RtMsghdr{}.Pid)) - 16]byte
	// RtMsghdr.Seq
	_ [20 - int(unsafe.Offsetof(RtMsghdr{}.Seq))]byte
	_ [int(unsafe.Offsetof(RtMsghdr{}.Seq)) - 20]byte
	// RtMsghdr.Errno
	_ [24 - int(unsafe.Offsetof(RtMsghdr{}.Errno))]byte
	_ [int(unsafe.Offsetof(RtMsghdr{}.Errno)) - 24]byte
	// RtMsghdr.Use
	_ [28 - int(unsafe.Offsetof(RtMsghdr{}.Use))]byte
	_ [int(unsafe.Offsetof(RtMsghdr{}.Use)) - 28]byte
	// RtMsghdr.Inits
	_ [32 - int(unsafe.Offsetof(RtMsghdr{}.Inits))]byte
	_ [int(unsafe.Offsetof(RtMsghdr{}.Inits)) - 32]byte
	// RtMsghdr.Rmx
	_ [36 - int(unsafe.Offsetof(RtMsghdr{}.Rmx))]byte
	_ [int(unsafe.Offsetof(RtMsghdr{}.Rmx)) - 36]byte
	// RtMsghdr2.size
	_ [92 - int(unsafe.Sizeof(RtMsghdr2{}))]byte
	_ [int(unsafe.Sizeof(RtMsghdr2{})) - 92]byte
	// RtMsghdr2.align
	_ [4 - int(unsafe.Alignof(RtMsghdr2{}))]byte
	_ [int(unsafe.Alignof(RtMsghdr2{})) - 4]byte
	// RtMsghdr2.Msglen
	_ [0 - int(unsafe.Offsetof(RtMsghdr2{}.Msglen))]byte
	_ [int(unsafe.Offsetof(RtMsghdr2{}.Msglen)) - 0]byte
	// RtMsghdr2.Version
	_ [2 - int(unsafe.Offsetof(RtMsghdr2{}.Version))]byte
	_ [int(unsafe.Offsetof(RtMsghdr2{}.Version)) - 2]byte
	// RtMsghdr2.Type
	_ [3 - int(unsafe.Offsetof(RtMsghdr2{}.Type))]byte
	_ [int(unsafe.Offsetof(RtMsghdr2{}.Type)) - 3]byte
	// RtMsghdr2.Index
	_ [4 - int(unsafe.Offsetof(RtMsghdr2{}.Index))]byte
	_ [int(unsafe.Offsetof(RtMsghdr2{}.Index)) - 4]byte
	// RtMsghdr2.Flags
	_ [8 - int(unsafe.Offsetof(RtMsghdr2{}.Flags))]byte
	_ [int(unsafe.Offsetof(RtMsghdr2{}.Flags)) - 8]byte
	// RtMsghdr2.Addrs
	_ [12 - int(unsafe.Offsetof(RtMsghdr2{}.Addrs))]byte
	_ [int(unsafe.Offsetof(RtMsghdr2{}.Addrs)) - 12]byte
	// RtMsghdr2.Refcnt
	_ [16 - int(unsafe.Offsetof(RtMsghdr2{}.Refcnt))]byte
	_ [int(unsafe.Offsetof(RtMsghdr2{}.Refcnt)) - 16]byte
	// RtMsghdr2.Parentflags
	_ [20 - int(unsafe.Offsetof(RtMsghdr2{}.Parentflags))]byte
	_ [int(unsafe.Offsetof(RtMsghdr2{}.Parentflags)) - 20]byte
	// RtMsghdr2.Reserved
	_ [24 - int(unsafe.Offsetof(RtMsghdr2{}.Reserved))]byte
	_ [int(unsafe.Offsetof(RtMsghdr2{}.Reserved)) - 24]byte
	// RtMsghdr2.Use
	_ [28 - int(unsafe.Offsetof(RtMsghdr2{}.Use))]byte
	_ [int(unsafe.Offsetof(RtMsghdr2{}.Use)) - 28]byte
	// RtMsghdr2.Inits
	_ [32 - int(unsafe.Offsetof(RtMsghdr2{}.Inits))]byte
	_ [int(unsafe.Offsetof(RtMsghdr2{}.Inits)) - 32]byte
	// RtMsghdr2.Rmx
	_ [36 - int(unsafe.Offsetof(RtMsghdr2{}.Rmx))]byte
	_ [int(unsafe.Offsetof(RtMsghdr2{}.Rmx)) - 36]byte
	// RtMetrics.size
	_ [56 - int(unsafe.Sizeof(RtMetrics{}))]byte
	_ [int(unsafe.Sizeof(RtMetrics{})) - 56]byte
	// RtMetrics.align
	_ [4 - int(unsafe.Alignof(RtMetrics{}))]byte
	_ [int(unsafe.Alignof(RtMetrics{})) - 4]byte
	// RtMetrics.Locks
	_ [0 - int(unsafe.Offsetof(RtMetrics{}.Locks))]byte
	_ [int(unsafe.Offsetof(RtMetrics{}.Locks)) - 0]byte
	// RtMetrics.Mtu
	_ [4 - int(unsafe.Offsetof(RtMetrics{}.Mtu))]byte
	_ [int(unsafe.Offsetof(RtMetrics{}.Mtu)) - 4]byte
	// RtMetrics.Hopcount
	_ [8 - int(unsafe.Offsetof(RtMetrics{}.Hopcount))]byte
	_ [int(unsafe.Offsetof(RtMetrics{}.Hopcount)) - 8]byte
	// RtMetrics.Expire
	_ [12 - int(unsafe.Offsetof(RtMetrics{}.Expire))]byte
	_ [int(unsafe.Offsetof(RtMetrics{}.Expire)) - 12]byte
	// RtMetrics.Recvpipe
	_ [16 - int(unsafe.Offsetof(RtMetrics{}.Recvpipe))]byte
	_ [int(unsafe.Offsetof(RtMetrics{}.Recvpipe)) - 16]byte
	// RtMetrics.Sendpipe
	_ [20 - int(unsafe.Offsetof(RtMetrics{}.Sendpipe))]byte
	_ [int(unsafe.Offsetof(RtMetrics{}.Sendpipe)) - 20]byte
	// RtMetrics.Ssthresh
	_ [24 - int(unsafe.Offsetof(RtMetrics{}.Ssthresh))]byte
	_ [int(unsafe.Offsetof(RtMetrics{}.Ssthresh)) - 24]byte
	// RtMetrics.Rtt
	_ [28 - int(unsafe.Offsetof(RtMetrics{}.Rtt))]byte
	_ [int(unsafe.Offsetof(RtMetrics{}.Rtt)) - 28]byte
	// RtMetrics.Rttvar
	_ [32 - int(unsafe.Offsetof(RtMetrics{}.Rttvar))]byte
	_ [int(unsafe.Offsetof(RtMetrics{}.Rttvar)) - 32]byte
	// RtMetrics.Pksent
	_ [36 - int(unsafe.Offsetof(RtMetrics{}.Pksent))]byte
	_ [int(unsafe.Offsetof(RtMetrics{}.Pksent)) - 36]byte
	// RtMetrics.Filler
	_ [40 - int(unsafe.Offsetof(RtMetrics{}.Filler))]byte
	_ [int(unsafe.Offsetof(RtMetrics{}.Filler)) - 40]byte
	// IfMsghdr.embedded
	_ [4 - int(unsafe.Offsetof(struct {
		Lead  byte
		Value IfMsghdr
	}{}.Value))]byte
	_ [int(unsafe.Offsetof(struct {
		Lead  byte
		Value IfMsghdr
	}{}.Value)) - 4]byte
	// IfMsghdr2.embedded
	_ [8 - int(unsafe.Offsetof(struct {
		Lead  byte
		Value IfMsghdr2
	}{}.Value))]byte
	_ [int(unsafe.Offsetof(struct {
		Lead  byte
		Value IfMsghdr2
	}{}.Value)) - 8]byte
	// IfData.embedded
	_ [4 - int(unsafe.Offsetof(struct {
		Lead  byte
		Value IfData
	}{}.Value))]byte
	_ [int(unsafe.Offsetof(struct {
		Lead  byte
		Value IfData
	}{}.Value)) - 4]byte
	// IfData64.embedded
	_ [8 - int(unsafe.Offsetof(struct {
		Lead  byte
		Value IfData64
	}{}.Value))]byte
	_ [int(unsafe.Offsetof(struct {
		Lead  byte
		Value IfData64
	}{}.Value)) - 8]byte
	// IfaMsghdr.embedded
	_ [4 - int(unsafe.Offsetof(struct {
		Lead  byte
		Value IfaMsghdr
	}{}.Value))]byte
	_ [int(unsafe.Offsetof(struct {
		Lead  byte
		Value IfaMsghdr
	}{}.Value)) - 4]byte
	// IfmaMsghdr.embedded
	_ [4 - int(unsafe.Offsetof(struct {
		Lead  byte
		Value IfmaMsghdr
	}{}.Value))]byte
	_ [int(unsafe.Offsetof(struct {
		Lead  byte
		Value IfmaMsghdr
	}{}.Value)) - 4]byte
	// IfmaMsghdr2.embedded
	_ [4 - int(unsafe.Offsetof(struct {
		Lead  byte
		Value IfmaMsghdr2
	}{}.Value))]byte
	_ [int(unsafe.Offsetof(struct {
		Lead  byte
		Value IfmaMsghdr2
	}{}.Value)) - 4]byte
	// RtMsghdr.embedded
	_ [4 - int(unsafe.Offsetof(struct {
		Lead  byte
		Value RtMsghdr
	}{}.Value))]byte
	_ [int(unsafe.Offsetof(struct {
		Lead  byte
		Value RtMsghdr
	}{}.Value)) - 4]byte
	// RtMsghdr2.embedded
	_ [4 - int(unsafe.Offsetof(struct {
		Lead  byte
		Value RtMsghdr2
	}{}.Value))]byte
	_ [int(unsafe.Offsetof(struct {
		Lead  byte
		Value RtMsghdr2
	}{}.Value)) - 4]byte
	// RtMetrics.embedded
	_ [4 - int(unsafe.Offsetof(struct {
		Lead  byte
		Value RtMetrics
	}{}.Value))]byte
	_ [int(unsafe.Offsetof(struct {
		Lead  byte
		Value RtMetrics
	}{}.Value)) - 4]byte
)

// Keep the exported wire-size constants tied to the checked Go types.
var (
	_ [SizeofIfMsghdr - int(unsafe.Sizeof(IfMsghdr{}))]byte
	_ [int(unsafe.Sizeof(IfMsghdr{})) - SizeofIfMsghdr]byte
	_ [SizeofIfMsghdr2 - int(unsafe.Sizeof(IfMsghdr2{}))]byte
	_ [int(unsafe.Sizeof(IfMsghdr2{})) - SizeofIfMsghdr2]byte
	_ [SizeofIfData - int(unsafe.Sizeof(IfData{}))]byte
	_ [int(unsafe.Sizeof(IfData{})) - SizeofIfData]byte
	_ [SizeofIfData64 - int(unsafe.Sizeof(IfData64{}))]byte
	_ [int(unsafe.Sizeof(IfData64{})) - SizeofIfData64]byte
	_ [SizeofIfaMsghdr - int(unsafe.Sizeof(IfaMsghdr{}))]byte
	_ [int(unsafe.Sizeof(IfaMsghdr{})) - SizeofIfaMsghdr]byte
	_ [SizeofIfmaMsghdr - int(unsafe.Sizeof(IfmaMsghdr{}))]byte
	_ [int(unsafe.Sizeof(IfmaMsghdr{})) - SizeofIfmaMsghdr]byte
	_ [SizeofIfmaMsghdr2 - int(unsafe.Sizeof(IfmaMsghdr2{}))]byte
	_ [int(unsafe.Sizeof(IfmaMsghdr2{})) - SizeofIfmaMsghdr2]byte
	_ [SizeofRtMsghdr - int(unsafe.Sizeof(RtMsghdr{}))]byte
	_ [int(unsafe.Sizeof(RtMsghdr{})) - SizeofRtMsghdr]byte
	_ [SizeofRtMsghdr2 - int(unsafe.Sizeof(RtMsghdr2{}))]byte
	_ [int(unsafe.Sizeof(RtMsghdr2{})) - SizeofRtMsghdr2]byte
	_ [SizeofRtMetrics - int(unsafe.Sizeof(RtMetrics{}))]byte
	_ [int(unsafe.Sizeof(RtMetrics{})) - SizeofRtMetrics]byte
)
