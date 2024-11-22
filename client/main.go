package main

import (
	"bufio"
	"flag"
	"net"
)

func main() {

	serverAddressPtr := flag.String("server", ":8080", "the server used for tunneling")
	flag.Parse()

	host := flag.Arg(0)
	if host == "" {
		PrintError("Please specify a host")
		return
	}

	serverAddress := *serverAddressPtr
	PrintInfo("Opening tunnel to " + serverAddress + "...")
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		PrintError("Failed to connect to " + serverAddress)
		return
	}
	defer conn.Close()

	PrintInfo("Tunnel opened. Forwarding information to " + host + ".")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		PrintInfo(scanner.Text())
	}
}
