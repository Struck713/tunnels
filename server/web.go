package server

import (
	"net"

	"nstruck.dev/tunnels/logger"
)

func InitWeb(address string) {

	logger.Info("Web", "Binding tunnel server to "+address)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.Error("Web failed to bind to " + address)
		return
	}
	defer listener.Close()

	logger.Info("Web", "Web server is now awaiting connections..")
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("Failed to establish web connection")
			continue
		}

		logger.Info("Web", "New connection from "+conn.RemoteAddr().String())
		go handleWeb(conn)
	}
}

func handleWeb(conn net.Conn) {
	defer conn.Close()
}
