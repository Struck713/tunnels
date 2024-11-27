package socket

import "net/http"

type HandshakeIdentity struct {
	Guid   string
	Domain string
}

type HandshakeAuthentication struct {
	Key string
}

type PageRequest struct {
	Headers http.Header
	URI     string
}

type PageResponse struct {
	Content string
	Headers http.Header
}
