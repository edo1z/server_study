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
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		fmt.Println("listen start")
		conn, err := listener.Accept()
		fmt.Println("listen now")
		if err != nil {
			fmt.Println("listen error")
			fmt.Println(err)
			continue
		}
		fmt.Println("listen ok")
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	fmt.Println("handle start")
	defer conn.Close()
	daytime := time.Now().String()
	conn.Write([]byte(daytime))
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)

	}
}
