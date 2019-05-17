package main

import (
	"code.byted.org/gopkg/logs"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

var host = flag.String("host", "", "host")
var port = flag.String("port", "3333", "port")

func main() {
	flag.Parse()
	var l net.Listener
	var err error
	l, err = net.Listen("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("error listening :", err)
		os.Exit(1)
	}
	defer l.Close()
	logs.Info("Listening host : %v , on port : %v", *host, *port)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("error accepting :", err)
			os.Exit(1)
		}
		logs.Info("Receved mgs %v -> %v \n", conn.RemoteAddr(), conn.LocalAddr())
		go handleRequest(conn)
	}

}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	for {
		io.Copy(conn, conn)
	}
}
