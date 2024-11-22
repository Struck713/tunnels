package main

import (
	"fmt"
	"net"
)

const ADDRESS = "localhost:8080"

func main() {

	// Listen for incoming connections
	listener, err := net.Listen("tcp", ADDRESS)
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
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	fmt.Fprint(conn, "This is a test\n")
}
