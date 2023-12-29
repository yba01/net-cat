package main

import (
	"os"
	"fmt"
	"net"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	CONN_PORT := "8989"
	if len(os.Args) == 2 {
		CONN_PORT = os.Args[1]
	}
	listener, err := net.Listen("tcp", "localhost:"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:",err.Error())
		return
	}
	defer listener.Close()
	fmt.Println("Listening on the port :"+CONN_PORT)
}