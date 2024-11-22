package socket

type HandshakeOutbound struct {
	Guid string
}

type PageRequestOutbound struct {
	Request string
}

type PageRequestInbound struct {
	Response string
}
