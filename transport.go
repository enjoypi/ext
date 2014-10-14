package ext

import (
	"net"
)

type Transport interface {
	Run()
}

type transport struct {
}

func NewTransport(net.Conn, Protocol) Transport {
	return &transport{}
}

func (t *transport) Run() {

}
