package http

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func Server() {
	listener, err := net.Listen("tcp", "localhost:8888")
	chkErr(err)
	fmt.Println("Server is running ato localhost:8888")
	for {
		conn, err := listener.Accept()
		chkErr(err)
		go func() {
			fmt.Printf("Accept %v\n", conn.RemoteAddr())
			request, err := http.ReadRequest(
				bufio.NewReader(conn),
			)
			chkErr(err)
			dump, err := httputil.DumpRequest(request, true)
			chkErr(err)
			fmt.Println(string(dump))
			response := http.Response{
				StatusCode: 200,
				ProtoMajor: 1,
				ProtoMinor: 0,
				Body: ioutil.NopCloser(
					strings.NewReader("Hello World\n"),
				),
			}
			_ = response.Write(conn)
			_ = conn.Close()
		}()
	}
}

func chkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(0)
	}
}
