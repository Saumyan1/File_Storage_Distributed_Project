package main

import (
	"fmt"
	"log"

	"github.com/Saumyan1/fileStorage/p2p"
)


func OnPeer(p2p.Peer) error{
	fmt.Println("doing some logiv with the peer outside of TCPtransport ")
	return nil
}

func main(){
	tcpOPts := p2p.TCPTransportOpts{
		ListenAddr: ":3000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder: p2p.DefaultDecoder{} ,
		OnPeer: OnPeer,


	}
	tr := p2p.NewTCPTransport(tcpOPts)
	if err := tr.ListenAndAccept(); err != nil{
		log.Fatal(err)
	}
	go func(){
		for{
			msg := <-tr.Consume()
			fmt.Printf("%v\n", msg)
		}
	}()

	select{}
	

}