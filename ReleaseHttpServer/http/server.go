package http

type HTTPServer struct {
	HTTPHandlers *HTTPHandlers
}

func NewHTTPServer(hand *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		HTTPHandlers: hand,
	}
}
