package server

import (
	"net"

	"github.com/google/uuid"
	"nstruck.dev/tunnels/logger"
	"nstruck.dev/tunnels/socket"
)

func InitServer(address string, web string, subdomain string) {
	clients := make(map[string]Client)
	go InitWeb(web, subdomain, clients)

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

		guid := uuid.New().String()
		logger.Info("Server", conn.RemoteAddr().String()+" <-> "+guid+"."+subdomain)
		socket.Send(conn, socket.HandshakeOutbound{
			Guid: guid,
			Domain:  guid + "." + subdomain,
		})

		client := Client{
			conn:     conn,
			request:  make(chan socket.PageRequest),
			response: make(chan socket.PageResponse),
		}
		clients[guid] = client
		go client.Init()
	}

}
