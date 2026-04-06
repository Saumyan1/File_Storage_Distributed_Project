package p2p



func NOPHandshakeFunc(Peer) error {return nil}


//Handshake func is
type HandshakeFunc func(Peer) error


