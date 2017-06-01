package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("connection failed!")
		return
	}
	defer conn.Close()

	conn.Write([]byte("hahaha"))
	bytes := make([]byte, 1024)
	conn.Read(bytes)
	fmt.Println(string(bytes))
}
