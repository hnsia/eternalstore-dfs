package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	tcpOpts := TCPTransportOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: NoopHandshakeFunc,
		Decoder:       DefaultDecoder{},
	}
	tr := NewTCPTransport(tcpOpts)
	assert.Equal(t, tcpOpts.ListenAddr, tr.ListenAddr)

	assert.Nil(t, tr.ListenAndAccept())
}
