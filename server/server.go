package server

import (
	"net"

	"github.com/google/uuid"
	"nstruck.dev/tunnels/logger"
	"nstruck.dev/tunnels/socket"
)

type Config struct {
	Address   string `json:"address"`
	Subdomain string `json:"subdomain"`
	Email     string `json:"email"`
}

func InitServer(config Config) {
	address := config.Address
	subdomain := config.Subdomain

	clients := make(map[string]Client)
	go InitWeb(subdomain, config.Email, clients)

	logger.Info("Server", "Binding tunnel server to "+address)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.Error("Server failed to bind to " + address)
		return
	}
	defer listener.Close()

	key := uuid.New().String()
	logger.Info("Server", "Authentication key: "+key)
	logger.Info("Server", "Tunnel server is now awaiting connections..")
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("Failed to establish tunnel connection")
			continue
		}

		packet := socket.Recieve[socket.HandshakeAuthentication](conn)
		if packet.Key != key {
			conn.Close()
			continue
		}

		guid := uuid.New().String()
		logger.Info("Server", conn.RemoteAddr().String()+" <-> "+guid+"."+subdomain)
		socket.Send(conn, socket.HandshakeIdentity{
			Guid:   guid,
			Domain: guid + "." + subdomain,
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
