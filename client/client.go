package client

import (
	"net"

	"nstruck.dev/tunnels/logger"
	"nstruck.dev/tunnels/socket"
)

func InitTunnel(from string, to string) {
	logger.Info("Client", "Opening tunnel to "+to+"...")
	conn, err := net.Dial("tcp", to)
	if err != nil {
		logger.Error("Failed to connect to " + to)
		return
	}
	defer conn.Close()

	handshake := socket.Recieve[socket.HandshakeOutbound](conn)
	if handshake == nil {
		logger.Error("Failed to handshake with tunnel.")
		return
	}
	logger.Info("Client", "Obtained unique ID from tunnel: " + handshake.Guid)
	logger.Info("Client", "Tunnel opened. Forwarding information to "+from+".")

}
