// Port-local Leopard capability declarations for darwin/ppc64 (goppc-p9apt.14).
// The sentinels and types below are ABSENT from Mac OS X 10.5.8 headers/kernel.
// No libc imports are made for the corresponding unavailable operations.
// Types marked derived are x/sys source-compatibility shapes only: they are
// never treated as Leopard ABI layouts, and unavailable calls fail explicitly.
//go:build darwin && ppc64

package unix

// ABSENT on Leopard: no AT_* constants or *at system calls. AT_FDCWD is a
// sentinel used only by the path-based compatibility wrappers below.
const (
	AT_FDCWD            = -0x2
	AT_REMOVEDIR        = 0x80
	AT_SYMLINK_FOLLOW   = 0x40
	AT_SYMLINK_NOFOLLOW = 0x20
	AT_EACCESS          = 0x10
	_AT_FDCWD           = AT_FDCWD
)

// F_DUPFD_CLOEXEC is absent from the Leopard 10.5.8 SDK. Keep the x/sys v0.48.0
// command value (zerrors_darwin_amd64.go:503), so old kernels reject it with
// EINVAL rather than treating zero as F_DUPFD and silently losing close-on-exec.
const F_DUPFD_CLOEXEC = 0x43

// SOL_LOCAL is the protocol-specific local socket-option level. Darwin defines
// it as zero (x/sys v0.48.0 zerrors_darwin_amd64.go:1373); XNU 10.5 retains
// LOCAL_PEERCRED at this level.
const SOL_LOCAL = 0

// ABSENT on Leopard: AF_VSOCK is not defined (AF_MAX is 37 on 10.5.8).
const AF_VSOCK = 0x28

// ABSENT on Leopard: no vsock address family. These source-compatibility
// values match x/sys v0.48.0 zerrors_darwin_amd64.go:1585-1589; Leopard rejects
// AF_VSOCK operations at the syscall.
const (
	VMADDR_CID_ANY        = 0xffffffff
	VMADDR_CID_HOST       = 0x2
	VMADDR_CID_HYPERVISOR = 0x0
	VMADDR_CID_RESERVED   = 0x1
	VMADDR_PORT_ANY       = 0xffffffff
)

// ABSENT on Leopard: no sys/vsock.h. This type is needed only so the shared
// SockaddrVM adapter typechecks; AF_VSOCK itself is rejected by the kernel.
// derived from darwin/amd64 LP64 (arm64 agrees); not hardware-verified at ppc64 — pmg5 is down.
type RawSockaddrVM struct {
	Len       uint8
	Family    uint8
	Reserved1 uint16
	Port      uint32
	Cid       uint32
}

const SizeofSockaddrVM = 0xc

// ABSENT on Leopard: connectx(2) and sae_*_t do not exist; connectx returns ENOSYS.
// derived from darwin/amd64 LP64 (arm64 agrees); not hardware-verified at ppc64 — pmg5 is down.
type SaeAssocID uint32

// derived from darwin/amd64 LP64 (arm64 agrees); not hardware-verified at ppc64 — pmg5 is down.
type SaeConnID uint32

// derived from darwin/amd64 LP64 (arm64 agrees); not hardware-verified at ppc64 — pmg5 is down.
type SaEndpoints struct {
	Srcif      uint32
	Srcaddr    *RawSockaddr
	Srcaddrlen uint32
	Dstaddr    *RawSockaddr
	Dstaddrlen uint32
	_          [4]byte
}

// ABSENT on Leopard: no struct ip_mreqn.
// derived from darwin/amd64 LP64 (arm64 agrees); not hardware-verified at ppc64 — pmg5 is down.
type IPMreqn struct {
	Multiaddr [4]byte /* in_addr */
	Address   [4]byte /* in_addr */
	Ifindex   int32
}

const SizeofIPMreqn = 0xc

// ABSENT on Leopard: no TCP_CONNECTION_INFO socket option.
// derived from darwin/amd64 LP64 (arm64 agrees); not hardware-verified at ppc64 — pmg5 is down.
type TCPConnectionInfo struct {
	State               uint8
	Snd_wscale          uint8
	Rcv_wscale          uint8
	_                   uint8
	Options             uint32
	Flags               uint32
	Rto                 uint32
	Maxseg              uint32
	Snd_ssthresh        uint32
	Snd_cwnd            uint32
	Snd_wnd             uint32
	Snd_sbbytes         uint32
	Rcv_wnd             uint32
	Rttcur              uint32
	Srtt                uint32
	Rttvar              uint32
	Txpackets           uint64
	Txbytes             uint64
	Txretransmitbytes   uint64
	Rxpackets           uint64
	Rxbytes             uint64
	Rxoutoforderbytes   uint64
	Txretransmitpackets uint64
}

const SizeofTCPConnectionInfo = 0x70
