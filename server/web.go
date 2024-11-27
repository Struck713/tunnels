package server

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/crypto/acme/autocert"
	"nstruck.dev/tunnels/logger"
	"nstruck.dev/tunnels/socket"
)

func InitWeb(address string, subdomain string, clients map[string]Client) {

	logger.Warning("Web", "Assigning autocerts to: *."+subdomain)

	allowedHosts := regexp.MustCompile("[^.]+." + subdomain)
	certManager := autocert.Manager{
		Prompt: autocert.AcceptTOS,
		HostPolicy: func(_ context.Context, host string) error {
			if matches := allowedHosts.MatchString(host); !matches {
				return errors.New("the host did not match the allowed hosts")
			}
			return nil
		},
		Cache: autocert.DirCache("certs"),
	}

	logger.Warning("Web", "Binding web server to "+address)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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

	server := &http.Server{
		Addr: ":https",
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
			MinVersion:     tls.VersionTLS12,
		},
	}

	go http.ListenAndServe(address, certManager.HTTPHandler(nil))
	logger.Error(server.ListenAndServeTLS("", "").Error())
}
