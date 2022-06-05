package server

import (
	"net/http"

	"htdvisser.dev/exp/connect-lambda/api/gen/go/htdvisser/echo/v1/echov1connect"
	"htdvisser.dev/exp/connect-lambda/internal/echoserver"
)

type Server struct {
	Mux *http.ServeMux
}

func NewServer() *Server {
	s := &Server{
		Mux: http.NewServeMux(),
	}

	s.Mux.Handle(echov1connect.NewEchoServiceHandler(&echoserver.EchoV1Server{}))

	return s
}
