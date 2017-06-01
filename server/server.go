package main

import (
	"fmt"
	"net"
)

// HandleConn HandleConn
func HandleConn(conn net.Conn) {
	defer conn.Close()
	bytes := make([]byte, 1024)
	for {
		_, err := conn.Read(bytes)
		if err == nil {
			conn.Write(bytes)
		} else {
			break
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error: ", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error: ", err)
			break
		}

		go HandleConn(conn)
	}
}
