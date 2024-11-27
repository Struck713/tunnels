package server

import (
	"net/http"
	"strings"

	"github.com/caddyserver/certmagic"
	"nstruck.dev/tunnels/logger"
	"nstruck.dev/tunnels/socket"
)

func InitWeb(subdomain string, email string, clients map[string]Client) {

	logger.Warning("Web", "Assigning autocerts to: *."+subdomain)
	logger.Warning("Web", "Binding web server to 0.0.0.0:80")

	mux := http.NewServeMux()

	certmagic.DefaultACME.Agreed = true
	certmagic.DefaultACME.Email = email
	certmagic.DefaultACME.CA = certmagic.LetsEncryptStagingCA

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		subdomains := strings.Split(r.Host, ".")
		domain := strings.Join(subdomains[1:], ".")
		if domain != subdomain {
			w.WriteHeader(404)
			return
		}

		client, exists := clients[subdomains[0]]
		if !exists {
			w.WriteHeader(404)
			return
		}

		client.request <- socket.PageRequest{
			URI:     r.RequestURI[1:],
			Headers: r.Header,
		}
		packet := <-client.response
		for key := range packet.Headers {
			w.Header().Add(key, packet.Headers.Get(key))
		}
		w.Write([]byte(packet.Content))
	})

	logger.Warning("Web", "Web server is now awaiting connections..")
	certmagic.HTTPS([]string{"*." + subdomain}, mux)
}
