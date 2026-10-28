//go:build darwin && ppc

package unix

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
