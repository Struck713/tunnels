package main

import (
	"flag"

	"nstruck.dev/tunnels/client"
	"nstruck.dev/tunnels/logger"
	"nstruck.dev/tunnels/server"
)

func main() {

	service := flag.String("service", "http://localhost:8085", "the service to tunnel")
	host := flag.String("hosts", "localhost:8081", "the tunnel server")
	// web := flag.String("web", "localhost:8083", "the web server")
	subdomain := flag.String("subdomain", "proxy.example.org", "the subdomain for the web server")
	email := flag.String("emails", "mail@proxy.example.org", "the email for ACME")
	flag.Parse()

	switch flag.Arg(0) {
	case "client":
		client.InitClient(flag.Arg(1), *service, *host)
		break
	case "server":
		server.InitServer(*host, *subdomain, *email)
		break
	default:
		logger.Error("Please either specify 'client' or 'server'")
		break
	}
}
