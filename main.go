package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	chkErr(err, "ResolveTCPAddr")

	listener, err := net.ListenTCP("tcp", tcpAddr)
	chkErr(err, "ListenTCP")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		daytime := time.Now().String()
		_, err = conn.Write([]byte(daytime))
		chkErr(err, "Write")

		err = conn.Close()
		chkErr(err, "Close")
	}

}

func chkErr(err error, place string) {
	if err != nil {
		fmt.Printf("(%s)", place)
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(0)
	}
}
