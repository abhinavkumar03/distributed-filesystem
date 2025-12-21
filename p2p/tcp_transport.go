package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPTransport is a Transport implementation that uses TCP for peer-to-peer communication.
type TCPTransport struct {
	listenAdderess string
	listener       net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		listenAdderess: listenAddress,
		peers:          make(map[net.Addr]Peer),
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAdderess)
	if err != nil {
		return err
	}
	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
			continue
		}
		go t.handleConnection(conn)

	}
}

func (t *TCPTransport) handleConnection(conn net.Conn) {
	fmt.Printf("New incoming connection from %+v\n", conn)
}
