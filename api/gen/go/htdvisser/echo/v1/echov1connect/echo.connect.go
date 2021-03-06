// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: htdvisser/echo/v1/echo.proto

package echov1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "htdvisser.dev/exp/connect-lambda/api/gen/go/htdvisser/echo/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// EchoServiceName is the fully-qualified name of the EchoService service.
	EchoServiceName = "htdvisser.echo.v1.EchoService"
)

// EchoServiceClient is a client for the htdvisser.echo.v1.EchoService service.
type EchoServiceClient interface {
	Echo(context.Context, *connect_go.Request[v1.EchoRequest]) (*connect_go.Response[v1.EchoResponse], error)
}

// NewEchoServiceClient constructs a client for the htdvisser.echo.v1.EchoService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewEchoServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) EchoServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &echoServiceClient{
		echo: connect_go.NewClient[v1.EchoRequest, v1.EchoResponse](
			httpClient,
			baseURL+"/htdvisser.echo.v1.EchoService/Echo",
			opts...,
		),
	}
}

// echoServiceClient implements EchoServiceClient.
type echoServiceClient struct {
	echo *connect_go.Client[v1.EchoRequest, v1.EchoResponse]
}

// Echo calls htdvisser.echo.v1.EchoService.Echo.
func (c *echoServiceClient) Echo(ctx context.Context, req *connect_go.Request[v1.EchoRequest]) (*connect_go.Response[v1.EchoResponse], error) {
	return c.echo.CallUnary(ctx, req)
}

// EchoServiceHandler is an implementation of the htdvisser.echo.v1.EchoService service.
type EchoServiceHandler interface {
	Echo(context.Context, *connect_go.Request[v1.EchoRequest]) (*connect_go.Response[v1.EchoResponse], error)
}

// NewEchoServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewEchoServiceHandler(svc EchoServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/htdvisser.echo.v1.EchoService/Echo", connect_go.NewUnaryHandler(
		"/htdvisser.echo.v1.EchoService/Echo",
		svc.Echo,
		opts...,
	))
	return "/htdvisser.echo.v1.EchoService/", mux
}

// UnimplementedEchoServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedEchoServiceHandler struct{}

func (UnimplementedEchoServiceHandler) Echo(context.Context, *connect_go.Request[v1.EchoRequest]) (*connect_go.Response[v1.EchoResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("htdvisser.echo.v1.EchoService.Echo is not implemented"))
}
