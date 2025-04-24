package Tools

import (
	"net"
	"sync"
)

type MessageStruct struct {
	Mesg []byte
	Conn net.Conn
	Infos []byte
	GenMessage []byte
}

var Chan = make(chan MessageStruct, 2)

var Clients = make(map[net.Conn]string)

var AllMessages string

var mu sync.Mutex