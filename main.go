package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"time"

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
		EncKey:            newEncryptionKey(),
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
	s2 := makeServer(":4000", ":3000") // Can only run in wsl or other unix OS, windows cannot create folders starting with ':'

	go func() {
		log.Fatal(s1.Start())
	}()
	time.Sleep(2 * time.Second)

	go s2.Start()
	time.Sleep(2 * time.Second)

	key := "coolPicture.jpg"
	data := bytes.NewReader([]byte("my big data file here!"))
	s2.Store(key, data)

	if err := s2.store.Delete(key); err != nil {
		log.Fatal(err)
	}

	r, err := s2.Get(key)
	// r, err := s2.Get("anewkeywedonthave")
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
