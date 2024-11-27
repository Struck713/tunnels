package client

import (
	"io"
	"net"
	"net/http"

	"nstruck.dev/tunnels/logger"
	"nstruck.dev/tunnels/socket"
)

func InitClient(key string, service string, server string) {
	logger.Info("Client", "Opening tunnel to "+server+"...")
	conn, err := net.Dial("tcp", server)
	if err != nil {
		logger.Error("Failed to connect to " + server)
		return
	}
	defer conn.Close()

	socket.Send(conn, socket.HandshakeAuthentication{
		Key: key,
	})

	handshake := socket.Recieve[socket.HandshakeIdentity](conn)
	if handshake == nil {
		logger.Error("Failed to handshake with tunnel.")
		return
	}

	domain := handshake.Domain
	logger.Info("Client", "URL: "+domain)
	logger.Info("Client", "Tunnel opened <-> "+service)

	for {
		packet := socket.Recieve[socket.PageRequest](conn)
		url := service + "/" + packet.URI
		logger.Info("Client", url+" <-> "+domain+"/"+packet.URI)

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
