package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	"htdvisser.dev/exp/connect-lambda/internal/server"
)

var listen = flag.String("listen", "127.0.0.1:3000", "Listen address")

func main() {
	flag.Parse()

	s := server.NewServer()

	lis, err := net.Listen("tcp", *listen)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Listening on http://%s", lis.Addr().String())

	log.Print(http.Serve(lis, s.Mux))
}
