package main

import (
	"context"
	"fmt"
	"io"
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

	stream, err := client.CreateALotOfContents(context.Background())
	if err != nil {
		log.Fatalf("Cannot stream to the server: %v", err)
	}
	log.Print("Streaming...")
	waitc := make(chan struct{})
	go func() {
		for {
			content, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive content: %v", err)
			}
			log.Printf("Got response from server: id: %d, text: %s", content.Id, content.Text)
		}
	}()
	for i:=0; i<1000; i++ {
		stream.Send(&pb.NewContent{Text: fmt.Sprintf("Send with stream 🌊 #%d", i)})
	}

	stream.CloseSend()
	<-waitc
}
