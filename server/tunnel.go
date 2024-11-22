package server

import (
	"fmt"
	"net"
)

func InitTunnel(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		go handleTunnel(conn)
	}
}

func handleTunnel(conn net.Conn) {
	defer conn.Close()
	fmt.Fprint(conn, "This is a test\n")
}
