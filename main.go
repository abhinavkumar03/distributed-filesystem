package main

import (
	"log"

	"github.com/abhinavkumar03/distributed-filesystem/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAdder:    ":3000",
		HandshakerFunc: p2p.NOPHandshakeFunc,
		Decoder:        p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}

// cmd for tcp connection in git bash
// curl telnet://localhost:3000
