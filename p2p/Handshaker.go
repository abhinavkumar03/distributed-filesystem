package p2p

// HandshakeFunc... ?
type HandshakerFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
