package main

import (
	"flag"

	"nstruck.dev/tunnels/client"
	"nstruck.dev/tunnels/logger"
	"nstruck.dev/tunnels/server"
)

func main() {

	from := flag.String("from", ":8085", "the server used for tunneling")
	to := flag.String("to", ":8081", "the service to tunnel")
	flag.Parse()

	switch flag.Arg(0) {
	case "client":
		client.InitTunnel(*from, *to)
		break
	case "server":
		server.InitServer()
		break
	default:
		logger.Error("Please either specify 'client' or 'server'")
		break
	}
}
