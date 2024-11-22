package server

import (
	"os"
	"os/signal"
	"syscall"

	"nstruck.dev/tunnels/logger"
)

func InitServer() {
	go InitWeb("localhost:8080")
	go InitTunnel("localhost:8081")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	logger.Info("Server", "Server closed")
}
