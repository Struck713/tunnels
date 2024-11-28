package server

import (
	"net/http"
	"strings"

	"github.com/caddyserver/certmagic"
	"github.com/libdns/cloudflare"
	"go.uber.org/zap"
	"nstruck.dev/tunnels/logger"
	"nstruck.dev/tunnels/socket"
)

func InitWeb(config Config, clients map[string]Client) {

	subdomain := config.Subdomain

	logger.Info("Web", "Assigning autocerts to: *."+subdomain)
	logger.Info("Web", "Binding web server to 0.0.0.0:80")

	mux := http.NewServeMux()

	certmagic.DefaultACME.Agreed = true
	certmagic.DefaultACME.Email = config.Email
	certmagic.DefaultACME.CA = certmagic.LetsEncryptStagingCA
	certmagic.Default.Logger, _ = zap.Config{
		Encoding: "console",
	}.Build()

	certmagic.DefaultACME.DNS01Solver = &certmagic.DNS01Solver{
		DNSManager: certmagic.DNSManager{
			DNSProvider: &cloudflare.Provider{
				APIToken: config.CloudflareKey,
			},
		},
	}

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

		logger.Info("Web", "Request: "+r.Host+" -> "+client.conn.RemoteAddr().String())

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

	logger.Info("Web", "Web server is now awaiting connections..")
	certmagic.HTTPS([]string{"*." + subdomain}, mux)
}
