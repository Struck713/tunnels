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

	logger.Info("Client", "GUID: "+handshake.Guid)
	logger.Info("Client", "Tunnel opened <-> "+from)

	for {
		packet := socket.Recieve[socket.PageRequest](conn)
		url := from + "/" + packet.URI
		logger.Info("Client", "Making request to service: "+url)

		req, err := http.NewRequest("GET", url, nil)
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

		socket.Send(conn, socket.PageResponse{
			Content: string(content),
			Headers: resp.Header,
		})
	}
}
