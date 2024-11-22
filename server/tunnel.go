package server

import (
	"net"

	"github.com/google/uuid"
	"nstruck.dev/tunnels/logger"
	"nstruck.dev/tunnels/socket"
)

func InitTunnel(address string) {

	logger.Info("Tunnel", "Binding tunnel server to " + address)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.Error("Tunnel failed to bind to " + address)
		return
	}
	defer listener.Close()

	logger.Info("Tunnel", "Tunnel server is now awaiting connections..")
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("Failed to establish tunnel connection")
			continue
		}

		logger.Info("Tunnel", "New connection from " + conn.RemoteAddr().String())
		go handleTunnel(conn)
	}
}

func handleTunnel(conn net.Conn) {
	defer conn.Close()

	guid := uuid.New().String()
	socket.Send(conn, socket.HandshakeOutbound{
		Guid: guid,
	})
	logger.Info("Tunnel", "Issued unique ID to client: " + guid)

}
