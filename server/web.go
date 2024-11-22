package server

import (
	"fmt"
	"net"
)

func InitWeb(address string) {
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

		fmt.Println("New client: " + conn.RemoteAddr().String())
		go handleWeb(conn)
	}

}

func handleWeb(conn net.Conn) {
	defer conn.Close()
	fmt.Fprint(conn, "This is a test\n")
}
