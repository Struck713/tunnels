package client

import (
	"io"
	"net"
	"net/http"

	"nstruck.dev/tunnels/logger"
	"nstruck.dev/tunnels/socket"
)

func InitClient(from string, to string) {
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

	logger.Info("Client", "Obtained unique ID from tunnel: "+handshake.Guid)
	logger.Info("Client", "Tunnel opened. Forwarding information from "+from+".")

	for {
		packet := socket.Recieve[socket.PageRequestOutbound](conn)
		logger.Info("Client", "Making request to service: "+from+packet.Request)

		req, err := http.NewRequest("GET", from+packet.Request, nil)
		if err != nil {
			logger.Error("Failed to create request to service.")
			continue
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			logger.Error("Failed to make request to service")
			continue
		}

		defer resp.Body.Close()

		content, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.Error("Failed to read response body.")
			continue
		}

		socket.Send(conn, socket.PageRequestInbound{
			Content: string(content),
			Headers:  resp.Header,
		})
	}
}
