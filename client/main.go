package main

import (
	"context"
	"log"
	"time"

	pb "github.com/rizafahmi/grpc-demo/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9000"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		log.Fatalf("Cannot connect to server, %v", err)
	}

	defer conn.Close()
	client := pb.NewContentManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.CreateContent(ctx, &pb.NewContent{Text: "New content from 🐹"})

	if err != nil {
		log.Fatalf("Could not create content: %v", err)
	}

	log.Printf(`Content details:
	Text: %s
	Id: %d`, response.GetText(), response.GetId())
}
