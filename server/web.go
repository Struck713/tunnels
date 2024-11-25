package server

import (
	"net/http"

	"nstruck.dev/tunnels/logger"
	"nstruck.dev/tunnels/socket"
)

func InitWeb(address string, channel chan<- http.Request, inbound <-chan socket.PageRequestInbound) {
	logger.Warning("Web", "Binding web server to "+address)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		channel <- *r
		packet := <-inbound
		for key := range packet.Headers {
			w.Header().Add(key, packet.Headers.Get(key))
		}
		w.Write([]byte(packet.Content))
	})
	logger.Warning("Web", "Web server is now awaiting connections..")
	http.ListenAndServe(address, nil)
}
