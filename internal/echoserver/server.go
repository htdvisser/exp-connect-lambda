package echoserver

import (
	"context"
	"log"

	"github.com/bufbuild/connect-go"
	echov1 "htdvisser.dev/exp/connect-lambda/api/gen/go/htdvisser/echo/v1"
)

type EchoV1Server struct{}

func (*EchoV1Server) Echo(
	_ context.Context, req *connect.Request[echov1.EchoRequest],
) (*connect.Response[echov1.EchoResponse], error) {
	log.Println("Request headers: ", req.Header())
	log.Println("Request message: ", req.Msg.Message)

	res := connect.NewResponse(&echov1.EchoResponse{
		Message: req.Msg.Message,
	})
	res.Header().Set("X-Echo-Version", "v1")

	return res, nil
}
