package chat

import (
	"../util"
	"fmt"
)

func clientSender(client *ClientChat) {
	for {
		select {
		case buf := <-client.IN:
			client.Con.Write([]byte(buf))
		case <-client.Quit:
			client.Con.Close()
			break
		}
	}
}

func clientHandling(conn *net.TCPConn, ch chan string, lst *list.List) {
	buf := make([]byte, 1024)
	con.Read(buf)
	name := string(buf)
	newClient := &ClientChat{name, make(chan string), ch, conn, make(chan bool), lst}

	go clientSender(newClient)
	go clientReceiver(newClient)
	lst.PushBack(*newClient)
	ch <- name + " has join the chat"
}

func Server() {
	clientList := list.New()
	in := make(chan string)
	go handlingINOUT(in, clientList)

	ln, err := net.Lieten("tcp", ":7777")
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		go clientHandling(&conn, in, clientList)
	}

}
