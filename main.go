package main

import (
	
	"log"

	"github.com/Saumyan1/fileStorage/p2p"
)


func main(){
	tcpOPts := p2p.TCPTransportOpts{
		ListenAddr: ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder: p2p.DefaultDecoder{} ,


	}
	tr := p2p.NewTCPTransport(tcpOPts)
	if err := tr.ListenAndAccept(); err != nil{
		log.Fatal(err)
	}

	select{}
	

}