package socket

import "net/http"

type HandshakeOutbound struct {
	Guid string
}

type PageRequestOutbound struct {
	Request string
}

type PageRequestInbound struct {
	Content string
	Headers http.Header
}
