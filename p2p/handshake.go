package p2p

// HandshakeFunc is ...
type HandshakeFunc func(any) error

func NoopHandshakeFunc(any) error { return nil }
