package main

import (
	"encoding/json"
	"flag"
	"os"

	"nstruck.dev/tunnels/client"
	"nstruck.dev/tunnels/logger"
	"nstruck.dev/tunnels/server"
)

func main() {

	configFile := flag.String("config", "config.json", "The location of the config you want to use.")

	flag.Parse()
	switch flag.Arg(0) {
	case "client":
		conf := loadConfig[client.Config](*configFile)
		client.InitClient(flag.Arg(1), conf)
		break
	case "server":
		conf := loadConfig[server.Config](*configFile)
		server.InitServer(conf)
		break
	default:
		logger.Error("Please either specify 'client' or 'server'")
		break
	}
}

func loadConfig[T interface{}](file string) T {
	conf, err := os.ReadFile(file)
	if err != nil {
		logger.Panic("Failed to open file: " + file)
	}

	obj := new(T)
	if json.Unmarshal(conf, obj) != nil {
		logger.Panic("Failed to load config: " + file)
	}

	return *obj
}
