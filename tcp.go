// Deprecated: This package has moved into go-libp2p as a sub-package: github.com/libp2p/go-libp2p/p2p/transport/tcp.
package tcp

import (
	"time"

	"github.com/MultiverseChronicles/go-libp2p/core/network"
	"github.com/MultiverseChronicles/go-libp2p/core/transport"

	"github.com/MultiverseChronicles/go-libp2p/p2p/transport/tcp"
)

// Deprecated: use github.com/libp2p/go-libp2p/p2p/transport/tcp.Option instead.
type Option = tcp.Option

// Deprecated: use github.com/libp2p/go-libp2p/p2p/transport/tcp.DisableReuseport instead.
func DisableReuseport() Option {
	return tcp.DisableReuseport()
}

// Deprecated: use github.com/libp2p/go-libp2p/p2p/transport/tcp.WithConnectionTimeout instead.
func WithConnectionTimeout(d time.Duration) Option {
	return tcp.WithConnectionTimeout(d)
}

// TcpTransport is the TCP transport.
// TcpTransport is the TCP transport.
type TcpTransport struct {
	// Connection upgrader for upgrading insecure stream connections to
	// secure multiplex connections.
	upgrader transport.Upgrader

	// optional custom dialer to use for dialing. If set, it will be the *ONLY* dialer
	// used. The transport will not attempt to reuse the listen port to
	// dial or the shared TCP transport for dialing.
	overrideDialerForAddr DialerForAddr

	disableReuseport bool // Explicitly disable reuseport.
	enableMetrics    bool

	// share and demultiplex TCP listeners across multiple transports
	sharedTcp *tcpreuse.ConnMgr

	// TCP connect timeout
	connectTimeout time.Duration

	rcmgr network.ResourceManager

	reuse reuseport.Transport

	metricsCollector *aggregatingCollector
}

var _ transport.Transport = &TcpTransport{}
var _ transport.DialUpdater = &TcpTransport{}

// NewTCPTransport creates a tcp transport object that tracks dialers and listeners
// created. It represents an entire TCP stack (though it might not necessarily be).
// Deprecated: use github.com/libp2p/go-libp2p/p2p/transport/tcp.NewTCPTransport instead.
/* 
func NewTCPTransport(upgrader transport.Upgrader, rcmgr network.ResourceManager, opts ...Option) (*TcpTransport, error) {
	return tcp.NewTCPTransport(upgrader, rcmgr, opts...)
}
 */

// NewTCPTransport creates a tcp transport object that tracks dialers and listeners
// created.
func NewTCPTransport(upgrader transport.Upgrader, rcmgr network.ResourceManager, sharedTCP *tcpreuse.ConnMgr, opts ...Option) (*TcpTransport, error) {
	if rcmgr == nil {
		rcmgr = &network.NullResourceManager{}
	}
	tr := &TcpTransport{
		upgrader:       upgrader,
		connectTimeout: defaultConnectTimeout, // can be set by using the WithConnectionTimeout option
		rcmgr:          rcmgr,
		sharedTcp:      sharedTCP,
	}
	for _, o := range opts {
		if err := o(tr); err != nil {
			return nil, err
		}
	}
	return tr, nil
}
