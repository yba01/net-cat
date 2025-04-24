package main

import (
	"fmt"
	"log"
	"net"
	"netcat/Tools"
	"os"
	"time"
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
		log.Fatal(err)
	}
	defer listener.Close()
	fmt.Println("Listening on the port :" + CONN_PORT)

	Clients := make(map[net.Conn]string)

	for {
		if len(Clients) == 10 {
			continue
		}
		Conn, err := listener.Accept()

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		Conn.Write(Tools.Connected())
		nameclient := Tools.GiveName(Conn, Clients)

		for conn := range Clients {
			conn.Write([]byte("\n" + nameclient + " has joined our chat...\n"))
			conn.Write([]byte("[" + time.Now().String()[:19] + "]" + "[" + Clients[conn] + "]:"))
		}
		Clients[Conn] = nameclient
		msg := make(chan string)
		go Tools.ChatReadMessage(Conn, Clients, msg)
		go Tools.ChatWriteMessage(Conn, Clients, msg)
	}
}
