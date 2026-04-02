package p2p


//peer is an interface that represent the remote node
type Peer interface {

}

//transport is anything that handle communication
//between the nodes in network.This can be of the
//form(TCP,UDP, websockets,..)
type Transport interface{
	ListenAndAccept() error
}

