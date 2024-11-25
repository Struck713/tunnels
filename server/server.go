package server

import (
	"net"
	"net/http"

	"github.com/google/uuid"
	"nstruck.dev/tunnels/logger"
	"nstruck.dev/tunnels/socket"
)

func InitServer(address string) {

	inbound := make(chan http.Request)
	defer close(inbound)

	outbound := make(chan socket.PageRequestInbound)
	defer close(outbound)

	go InitWeb("localhost:8083", inbound, outbound)

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
		go handleTunnel(conn, inbound, outbound)
	}

}

func handleTunnel(conn net.Conn, inbound <-chan http.Request, outbound chan<- socket.PageRequestInbound) {
	defer conn.Close()

	guid := uuid.New().String()
	socket.Send(conn, socket.HandshakeOutbound{
		Guid: guid,
	})
	logger.Info("Server", "Issued unique ID to "+conn.RemoteAddr().String()+": "+guid)
	for {
		request := <-inbound
		socket.Send(conn, socket.PageRequestOutbound{
			Request: request.RequestURI,
		})

		packet := socket.Recieve[socket.PageRequestInbound](conn)
		if packet == nil {
			logger.Error("Failed to recieve response from tunnel.")
			continue
		}
		outbound <- *packet
	}

}
