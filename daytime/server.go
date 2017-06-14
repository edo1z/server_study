package daytime

import (
	"fmt"
	"net"
	"os"
	"time"
)

func Server() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	chkErr(err)

	li, err := net.ListenTCP("tcp", tcpAddr)
	chkErr(err)

	for {
		conn, err := li.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		_, err = conn.Write([]byte(daytime))
		chkErr(err)
		_ = conn.Close()
	}
}

func chkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(0)
	}
}
