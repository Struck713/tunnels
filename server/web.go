package server

import (
	"net/http"
	"strings"

	"nstruck.dev/tunnels/logger"
	"nstruck.dev/tunnels/socket"
)

func InitWeb(address string, clients map[string]Client) {
	logger.Warning("Web", "Binding web server to "+address)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		paths := strings.Split(r.RequestURI[1:], "/")
		guid := paths[0]
		client, exists := clients[guid]
		if !exists {
			w.Write([]byte("No page found."))
			return
		}

		client.request <- socket.PageRequest{
			URI:     strings.Join(paths[1:], "/"),
			Headers: r.Header,
		}
		packet := <-client.response
		for key := range packet.Headers {
			w.Header().Add(key, packet.Headers.Get(key))
		}
		w.Write([]byte(packet.Content))
	})
	logger.Warning("Web", "Web server is now awaiting connections..")
	http.ListenAndServe(address, nil)
}
