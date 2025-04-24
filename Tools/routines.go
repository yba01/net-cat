package Tools

import (
	"net"
	"time"
)

func HandleConnections(conn net.Conn, ClientName []byte) {
	defer conn.Close()
	Mess := make([]byte, 2048)
	var Message []byte
	for {
		infos := []byte("[" + time.Now().String()[:19] + "]" + "[" + string(ClientName) + "]:")
		conn.Write(infos)
		n, err := conn.Read(Mess)
		if len(Mess) > 2048 {
			conn.Write([]byte("Message too long\n"))
			continue
		}
		if err != nil {
			Chan <- MessageStruct{
				Mesg:       []byte("\n" + string(ClientName) + " has left the chat..." + "\n"),
				Conn:       conn,
				Infos:      []byte("[" + time.Now().String()[:19] + "]" + "[" + string(ClientName) + "]:"),
				GenMessage: []byte("yes"),
			}
			mu.Lock()
			delete(Clients, conn)
			mu.Unlock()
			break
		}
		Message = append(Message, Mess[:n]...)
		Chan <- MessageStruct{
			Mesg:       Message,
			Conn:       conn,
			Infos:      []byte("[" + time.Now().String()[:19] + "]" + "[" + string(ClientName) + "]:"),
			GenMessage: []byte("nope"),
		}
		Message = nil
	}
}

func ChatWriter() {
	for {
		MsgStruct := <-Chan
		if string(MsgStruct.GenMessage) == "yes" {
			AllMessages += string(MsgStruct.Mesg)[1:]
		} else {
			AllMessages += string(MsgStruct.Infos) + string(MsgStruct.Mesg)
		}
		if string(MsgStruct.Mesg)[:len(string(MsgStruct.Mesg))-1] == "" {
			continue
		}
		mu.Lock()
		for conn, name := range Clients {
			if MsgStruct.Conn != conn {
				if string(MsgStruct.GenMessage) == "yes" {
					conn.Write([]byte(string(MsgStruct.Mesg)))
				} else {
					conn.Write([]byte("\n" + string(MsgStruct.Infos) + string(MsgStruct.Mesg)))
				}
				conn.Write([]byte(string(MsgStruct.Infos)[:21] + "[" + name + "]:"))
			}
		}
		mu.Unlock()
	}
}
