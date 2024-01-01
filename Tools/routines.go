package Tools

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func ChatReadMessage(conn net.Conn, Chats map[net.Conn]string, msg chan string) {
	defer conn.Close()
	Mess := make([]byte, 2048)
	var Message []byte

	for {
		n, err := conn.Read(Mess)
		if err != nil {
			msg <- string([]byte(Chats[conn] + " has left our chat...\n"))
			break
		}
		Message = append(Message, Mess[:n]...)
		if string(Message) != "\n" {
			msg <- string([]byte("["+time.Now().String()[:19]+"]"+"[" + Chats[conn]+ "]:" + string(Message)))
		}
		Message = nil
		conn.Write([]byte("["+time.Now().String()[:19]+"]"+"[" + Chats[conn]+ "]:"))
	}
	delete(Chats, conn)
}

func ChatWriteMessage(connec net.Conn, Chats map[net.Conn]string, msg chan string) {
	for {
		chatting := <-msg
		for conn := range Chats {
			if connec != conn {
				conn.Write([]byte("\n" + chatting))
				conn.Write([]byte("["+time.Now().String()[:19]+"]"+"["+Chats[conn]+"]:"))
			}
		}
	}
}
func Connected() []byte {
	file, err := os.Open("Chat/welcome.txt")

	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()

	var welcome []byte
	text := bufio.NewScanner(file)

	for text.Scan() {
		welcome = append(welcome, text.Bytes()...)
		welcome = append(welcome, '\n')
	}

	return welcome
}
func GiveName(Conn net.Conn, Clients map[net.Conn]string) string {
	ClientName := make([]byte, 1024)
	nameclient:=""
	for {
		arret := true
		Conn.Write([]byte("[ENTER YOUR NAME]:"))
		n,_:=Conn.Read(ClientName)
		nameclient = string(ClientName[:n-1])
		for _,name := range Clients {
			if nameclient== name {
				Conn.Write([]byte("NAME ALREADY TAKEN, CHOOSE ANOTHER NAME\n"))
				arret = false
			}
		}
		if arret && nameclient != "" {
			break
		}
	}
	Conn.Write([]byte("["+time.Now().String()[:19]+"]"+"["+nameclient+"]:"))
	return nameclient
}
