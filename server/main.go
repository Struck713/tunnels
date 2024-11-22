package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	conn.Close()
}
