package simple_chat

import (
	"../util"
	"fmt"
	"net"
	"time"
)

var running = true

func sender(conn net.Conn) {
}

func receiver(conn net.Conn) {
	for running {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		fmt.Println("received")
		util.ChkErr(err, "Receiver read")
		fmt.Println(string(buf[:n]))
	}
}

func Server() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	util.ChkErr(err, "tcpaddr")

	li, err := net.ListenTCP("tcp", tcpAddr)
	util.ChkErr(err, "tcpaddr")

	conn, err := li.Accept()
	util.ChkErr(err, "accept")

	go receiver(conn)
	go sender(conn)

	for running {
		time.Sleep(1 * 1e9)
	}
}
