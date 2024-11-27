package main

import (
	"flag"

	"nstruck.dev/tunnels/client"
	"nstruck.dev/tunnels/logger"
	"nstruck.dev/tunnels/server"
)

func main() {

	service := flag.String("service", "http://localhost:8085", "the service to tunnel")
	host := flag.String("host", "localhost:8081", "the tunnel server")
	web := flag.String("web", "localhost:8083", "the web server")
	subdomain := flag.String("subdomain", "proxy.nstruck.dev", "the subdomain for the web server")
	flag.Parse()

	switch flag.Arg(0) {
	case "client":
		client.InitClient(*service, *host)
		break
	case "server":
		server.InitServer(*host, *web, *subdomain)
		break
	default:
		logger.Error("Please either specify 'client' or 'server'")
		break
	}
}
