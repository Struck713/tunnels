package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func InitServer() {
	go InitWeb("localhost:8080")
	go InitTunnel("localhost:8081")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	fmt.Println("Shutting down..")
}
