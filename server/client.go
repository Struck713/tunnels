package server

import (
	"net"

	"nstruck.dev/tunnels/logger"
	"nstruck.dev/tunnels/socket"
)

type Client struct {
	conn     net.Conn
	request  chan socket.PageRequest
	response chan socket.PageResponse
}

func (c *Client) Init() {
	defer c.Close()
	for {
		requestPacket := <-c.request
		socket.Send(c.conn, requestPacket)

		responsePacket := socket.Recieve[socket.PageResponse](c.conn)
		if responsePacket == nil {
			logger.Error("Failed to recieve response from tunnel.")
			continue
		}
		c.response <- *responsePacket
	}
}

func (c *Client) Close() {
	close(c.request)
	close(c.request)
	c.conn.Close()
}
