package socket

import "net/http"

type HandshakeOutbound struct {
	Guid   string
	Domain string
}

type PageRequest struct {
	Headers http.Header
	URI     string
}

type PageResponse struct {
	Content string
	Headers http.Header
}
