package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"htdvisser.dev/exp/connect-lambda/internal/server"
)

func main() {
	s := server.NewServer()
	adapter := httpadapter.NewV2(s.Mux)
	lambda.Start(adapter.ProxyWithContext)
}
