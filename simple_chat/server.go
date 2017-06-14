package simple_chat

import (
	"../util"
	"fmt"
	"net"
	"time"
)

func sender(conn net.Conn) {

}

func receiver(conn net.Conn) {
	_ = conn.SetReadDeadline(time.Now().Add(2 * time.Minute))
	buf := make([]byte, 560)

	for {
		n, err := conn.Read(buf)
		util.ChkErr(err, "Receiver read")
		fmt.Println(string(buf[:n]))

		buf = make([]byte, 560)
	}
}

func Server() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	util.ChkErr(err, "tcpaddr")

	li, err := net.ListenTCP("tcp", tcpAddr)
	util.ChkErr(err, "tcpaddr")

	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println("Fail to connect.")
			continue
		}

		go receiver(conn)
		go sender(conn)
	}
}
