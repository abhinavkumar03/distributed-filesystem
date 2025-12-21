package p2p

// Peer is an interface that represents the remote node in the network.
type Peer interface{}

// Transport represents a transport mechanism for peer-to-peer communication.
// Transport is anything that handles the communication between the nodes in the network.
// This can be of the form (TCP, UDP, WebRTC, etc.)
type Transport interface {
	ListenAndAccept() error
}
