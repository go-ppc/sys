//go:build darwin && ppc64

package unix

// The same 21 supplemental constants the ppc32 layer carries
// (zconstants_leopard_supp_darwin_ppc.go), for darwin/ppc64. goppc-p9apt.35.
//
// Provenance, stated at the level the evidence supports:
//
//   - **These names do not exist in Mac OS X 10.5.8's C environment.** All 21 are
//     absent from `gcc -arch ppc -m32 -E -dM` *and* from `gcc -arch ppc64 -m64 -E -dM`
//     over the relevant headers on a real Leopard machine, and none appears in the
//     Leopard-era xnu tree in this workspace (`reference/xnu/bsd`):
//     runs/xsys-ppc64-cc-20261028/receipt.txt and
//     runs/xsys-ppc64-cc-on-g4-20261028/receipt.txt. So no C compiler on this
//     platform can be asked for their values; "Leopard supplemental" means
//     supplemental *to* Leopard, not present *in* it.
//   - **The values are width-independent by construction.** Each one equals the
//     value upstream `x/sys` declares for `darwin/amd64` *and* `darwin/arm64`
//     (21 of 21 checked by name against v0.48.0's `zerrors_darwin_amd64.go` and
//     `zerrors_darwin_arm64.go`), and each is a plain `int` option number, not a
//     size or an offset. Nothing here is derived from the 64-bit ABI.
//   - **Runtime behaviour on Leopard is not "success".** These are RFC 3542 /
//     RFC 2113-era options; the Leopard kernel rejects the socket ones
//     (`IP_*`, `IPV6_*`, `MCAST_*`) with an errno such as `ENOPROTOOPT`, which is
//     an explicit failure through the existing ABI, not a silent approximation.
//     `MNT_CPROTECT` is a mount-flag bit read out of `statfs`, so the only effect
//     of the constant is to test a bit the platform may never set.
//   - **Nothing in this file has been executed at ppc64.** No 64-bit PowerPC
//     machine is available to this project (`pmg5` is down); the ppc32 layer's
//     constants are exercised on a real G4, and this file mirrors it.
//
// value from golang.org/x/sys v0.48.0 zerrors_darwin_amd64.go:777
const IPV6_2292HOPLIMIT = 0x14

// value from golang.org/x/sys v0.48.0 zerrors_darwin_amd64.go:779
const IPV6_2292NEXTHOP = 0x15

// value from golang.org/x/sys v0.48.0 zerrors_darwin_amd64.go:780
const IPV6_2292PKTINFO = 0x13

// value from golang.org/x/sys v0.48.0 zerrors_darwin_amd64.go:799
const IPV6_DONTFRAG = 0x3e

// value from golang.org/x/sys v0.48.0 zerrors_darwin_amd64.go:830
const IPV6_PATHMTU = 0x2c

// value from golang.org/x/sys v0.48.0 zerrors_darwin_amd64.go:838
const IPV6_RECVHOPLIMIT = 0x25

// value from golang.org/x/sys v0.48.0 zerrors_darwin_amd64.go:840
const IPV6_RECVPATHMTU = 0x2b

// value from golang.org/x/sys v0.48.0 zerrors_darwin_amd64.go:841
const IPV6_RECVPKTINFO = 0x3d

// value from golang.org/x/sys v0.48.0 zerrors_darwin_amd64.go:843
const IPV6_RECVTCLASS = 0x23

// value from golang.org/x/sys v0.48.0 zerrors_darwin_amd64.go:850
const IPV6_TCLASS = 0x24

// value from golang.org/x/sys v0.48.0 zerrors_darwin_amd64.go:863
const IP_DONTFRAG = 0x1c

// value from golang.org/x/sys v0.48.0 zerrors_darwin_amd64.go:902
const IP_PKTINFO = 0x1a

// value from golang.org/x/sys v0.48.0 zerrors_darwin_amd64.go:910
const IP_RECVPKTINFO = 0x1a

// value from golang.org/x/sys v0.48.0 zerrors_darwin_amd64.go:912
const IP_RECVTOS = 0x1b

// value from golang.org/x/sys v0.48.0 zerrors_darwin_amd64.go:975
const MCAST_BLOCK_SOURCE = 0x54

// value from golang.org/x/sys v0.48.0 zerrors_darwin_amd64.go:978
const MCAST_JOIN_GROUP = 0x50

// value from golang.org/x/sys v0.48.0 zerrors_darwin_amd64.go:979
const MCAST_JOIN_SOURCE_GROUP = 0x52

// value from golang.org/x/sys v0.48.0 zerrors_darwin_amd64.go:980
const MCAST_LEAVE_GROUP = 0x51

// value from golang.org/x/sys v0.48.0 zerrors_darwin_amd64.go:981
const MCAST_LEAVE_SOURCE_GROUP = 0x53

// value from golang.org/x/sys v0.48.0 zerrors_darwin_amd64.go:982
const MCAST_UNBLOCK_SOURCE = 0x55

// value from golang.org/x/sys v0.48.0 zerrors_darwin_amd64.go:989
const MNT_CPROTECT = 0x80
