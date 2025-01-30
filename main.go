package main

import (
	"bytes"
	"log"

	"github.com/hnsia/eternalstore-dfs/p2p"
)

func OnPeer(peer p2p.Peer) error {
	// return fmt.Errorf("failed onpeer func")
	peer.Close()
	return nil
}

func makeServer(listenAddr string, nodes ...string) *FileServer {
	tcpTransportOpts := p2p.TCPTransportOpts{
		ListenAddr:    listenAddr,
		HandshakeFunc: p2p.NoopHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		// TODO: onPeer func
	}
	tcpTransport := p2p.NewTCPTransport(tcpTransportOpts)

	fileServerOpts := FileServerOpts{
		StorageRoot:       listenAddr + "_network",
		PathTransformFunc: CASPathTransformFunc,
		Transport:         tcpTransport,
		BootstrapNodes:    nodes,
	}

	s := NewFileServer(fileServerOpts)

	tcpTransport.OnPeer = s.OnPeer

	return s
}

func main() {
	s1 := makeServer(":3000", "")
	s2 := makeServer(":4000", ":3000")

	go func() {
		log.Fatal(s1.Start())
	}()

	s2.Start()

	data := bytes.NewReader([]byte("my big data file here!"))

	s2.StoreData("myprivatedata", data)
}
