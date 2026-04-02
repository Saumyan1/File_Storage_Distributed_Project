package p2p

import (
	"bytes"
	"fmt"
	"net"
	"sync"
)

//Im here building a transport layer that is responsible for
//sending and recieving data between machinesa

//TCPPeer represents the remote node over a TCP established connection
type TCPPeer struct{
	//conn is the underline connection of the peer
	conn net.Conn
	//if we dial and retrieve a conn => outbound == true
	//if we accept and retrieve a conn => outbound == false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer{
	return &TCPPeer{
		conn: conn,
		outbound: outbound,
	}

}

//This is a TCP server that listens on listenAddress,
//uses a listener to accept connections, and maintains
// a peers map to track all connected nodes
//transport always listens and accept
type TCPTransport struct {
	listneAddress string
	listner net.Listener
	shakeHands HandshakeFunc
	decoder Decoder
	mu sync.RWMutex
	peers map[net.Addr]Peer
}


func NewTCPTransport(listenAddr string) *TCPTransport{
	return &TCPTransport{
		shakeHands: NOPHandshakeFunc,
		listneAddress: listenAddr,
	}
}

func(t *TCPTransport) ListenAndAccept()error{
	var err error
	t.listner,err = net.Listen("tcp",t.listneAddress)
	if err != nil{
		
		return err
	}

	go t.startAcceptLoop()
	return nil


}


func (t *TCPTransport)startAcceptLoop(){

	//If i had used no for loop then it would
	//only accept 1 connection
	for{
		conn, err := t.listner.Accept()
		if err != nil{
			fmt.Printf("TCP accept error: %s\n",err)
		}

		go t.handleConn(conn)

	}

}

type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn){
	peer := NewTCPPeer(conn, true)

	if err := t.shakeHands(conn); err != nil{

	}

	

	//readloop
	msg:= &Temp{}
	for{
		if err := t.decoder.Decode(conn,msg); err != nil{
			fmt.Printf("TCP error: %s\n", err)
			continue
		}

	}

	fmt.Printf("new incoming connection%+v\n",peer)

}