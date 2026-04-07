package p2p

import (
	
	"fmt"
	"net"

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
//Close implement the Peer interface
func (p *TCPPeer) Close() error{
	return p.conn.Close()
}
type TCPTransportOpts struct{
	ListenAddr string
	HandshakeFunc HandshakeFunc
	Decoder Decoder
	OnPeer func(Peer) error
}

//This is a TCP server that listens on listenAddress,
//uses a listener to accept connections, and maintains
// a peers map to track all connected nodes
//transport always listens and accept
type TCPTransport struct {
	TCPTransportOpts
	listner net.Listener
	rpcchan chan	 RPC
}


func NewTCPTransport(opts TCPTransportOpts) *TCPTransport{
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpcchan: make(chan RPC),
	}
}

//consume implement the Transport interface, which will return read only
//channel for reading the incoming messgaes received from another peer
//in the network
func (t *TCPTransport) Consume() <-chan RPC{
	return t.rpcchan
}

func(t *TCPTransport) ListenAndAccept()error{
	var err error
	t.listner,err = net.Listen("tcp",t.ListenAddr)
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
		fmt.Printf("new incoming connection %+v\n",conn)

		go t.handleConn(conn)

	}

}

type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn){
	var err error
	defer func(){
		fmt.Printf("dropping peer connection: %s",err)
		conn.Close()
	}()
	peer := NewTCPPeer(conn, true)
	

	if err := t.HandshakeFunc(peer); err != nil{
		return

	}

	if t.OnPeer != nil{
		if err = t.OnPeer(peer); err != nil{
			return
		}
	}



	//readloop
	//here we are not decoding message but the RPC
	rpc:= RPC{}
	for{
		err = t.Decoder.Decode(conn,&rpc)
		if err != nil{
			fmt.Printf("TCP read error: %s\n", err)
			return
			
		}
		rpc.From = conn.RemoteAddr()
		t.rpcchan <- rpc

	}

}