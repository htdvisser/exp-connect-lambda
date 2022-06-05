package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/bufbuild/connect-go"
	echov1 "htdvisser.dev/exp/connect-lambda/api/gen/go/htdvisser/echo/v1"
	"htdvisser.dev/exp/connect-lambda/api/gen/go/htdvisser/echo/v1/echov1connect"
)

var (
	baseURL   = flag.String("base-url", "http://localhost:3000", "server base url")
	grpc      = flag.Bool("grpc", false, "use grpc protocol")
	grpcWeb   = flag.Bool("grpc-web", false, "use grpc-web protocol")
	protoJSON = flag.Bool("protojson", false, "use protojson encoding")
	authToken = flag.String("auth-token", "", "auth token")
)

func main() {
	flag.Parse()

	var clientOptions []connect.ClientOption

	if *grpc {
		clientOptions = append(clientOptions, connect.WithGRPC())
	}
	if *grpcWeb {
		clientOptions = append(clientOptions, connect.WithGRPCWeb())
	}
	if *protoJSON {
		clientOptions = append(clientOptions, connect.WithProtoJSON())
	}

	client := echov1connect.NewEchoServiceClient(
		http.DefaultClient,
		*baseURL,
		clientOptions...,
	)

	req := connect.NewRequest(&echov1.EchoRequest{
		Message: strings.Join(flag.Args(), " "),
	})
	if *authToken != "" {
		req.Header().Set("X-Auth-Token", *authToken)
	}

	res, err := client.Echo(context.Background(), req)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Response headers: ", res.Header())
	log.Println("Response message: ", res.Msg.Message)
}
