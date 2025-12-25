package main

import (
	"log"

	"github.com/abhinavkumar03/distributed-filesystem/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}

// cmd for tcp connection
// Test-NetConnection localhost -Port 3000
// curl http://localhost:3000/health
