package server

import (
	"net"

	"nstruck.dev/tunnels/logger"
)

func InitWeb(address string, channel chan TunnelRequest) {

	logger.Warning("Web", "Binding web server to "+address)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.Error("Web server failed to bind to " + address)
		return
	}
	defer listener.Close()

	logger.Warning("Web", "Web server is now awaiting connections..")
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("Failed to establish web connection")
			continue
		}

		defer conn.Close()
		channel<-TunnelRequest{
			Metadata: "request was recieved",
		}
	}
}
