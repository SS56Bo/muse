package p2p

import (
	"net"
	"sync"
)

// TCP transport to handle communications
// between p2p networks & nodes using TCP
type tcpTransportHandler struct {
	listener net.Listener   		// listens for incoming connections
	listenAddr string 	           //stores address
	mute sync.RWMutex             // read-write mutex
	peers map[net.Addr]Peer
}

func NewTCPConnection(listenAddress string) *tcpTransportHandler {
	return &tcpTransportHandler{
		listenAddr: listenAddress,
	}
}