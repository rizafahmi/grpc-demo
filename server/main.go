package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Cannot start the server %v", err)
	}

	server := grpc.NewServer()

	log.Printf("Server running at %v", listen.Addr())
	err = server.Serve(listen)

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
