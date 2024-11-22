package client

import (
	"bufio"
	"net"

	"nstruck.dev/tunnels/logger"
)

func InitTunnel(from string, to string) {
	logger.Info("Opening tunnel to " + to + "...")
	conn, err := net.Dial("tcp", to)
	if err != nil {
		logger.Error("Failed to connect to " + to)
		return
	}
	defer conn.Close()

	logger.Info("Tunnel opened. Forwarding information to " + from + ".")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		logger.Info(scanner.Text())
	}

}
