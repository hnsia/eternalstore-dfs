package main

import (
	"log"
	"time"

	"github.com/hnsia/eternalstore-dfs/p2p"
)

func OnPeer(peer p2p.Peer) error {
	// return fmt.Errorf("failed onpeer func")
	peer.Close()
	return nil
}

func main() {
	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: p2p.NoopHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		// TODO: onPeer func
	}
	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)

	fileServerOpts := FileServerOpts{
		StorageRoot:       "3000_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
	}

	s := NewFileServer(fileServerOpts)

	go func() {
		time.Sleep(time.Second * 3)
		s.Stop()
	}()

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

	// tcpOpts := p2p.TCPTransportOpts{
	// 	ListenAddr:    ":3000",
	// 	HandshakeFunc: p2p.NoopHandshakeFunc,
	// 	Decoder:       p2p.DefaultDecoder{},
	// 	OnPeer:        OnPeer,
	// }
	// tr := p2p.NewTCPTransport(tcpOpts)

	// go func() {
	// 	for {
	// 		msg := <-tr.Consume()
	// 		fmt.Printf("%+v\n", msg)
	// 	}
	// }()

	// if err := tr.ListenAndAccept(); err != nil {
	// 	log.Fatal(err)
	// }

	// select {}
}
