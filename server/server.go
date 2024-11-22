package server

import (
	"net"

	"github.com/google/uuid"
	"nstruck.dev/tunnels/logger"
	"nstruck.dev/tunnels/socket"
)

type TunnelRequest struct {
	Metadata string
}

func InitServer(address string) {

	channel := make(chan TunnelRequest)
	go InitWeb("localhost:8083", channel)

	logger.Info("Server", "Binding tunnel server to "+address)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.Error("Server failed to bind to " + address)
		return
	}
	defer listener.Close()

	logger.Info("Server", "Tunnel server is now awaiting connections..")
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("Failed to establish tunnel connection")
			continue
		}
		go handleTunnel(conn, channel)
	}
}

func handleTunnel(conn net.Conn, channel chan TunnelRequest) {
	defer conn.Close()

	guid := uuid.New().String()
	socket.Send(conn, socket.HandshakeOutbound{
		Guid: guid,
	})
	logger.Info("Server", "Issued unique ID to "+conn.RemoteAddr().String()+": "+guid)

	for {
		request := <-channel
		logger.Info("Server", "Web request passed: "+request.Metadata)
	}

}
