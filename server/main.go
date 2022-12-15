package main

import (
	"context"
	"io"
	"log"
	"math/rand"
	"net"

	pb "github.com/rizafahmi/grpc-demo/proto"
	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

type ContentManagementServer struct {
	pb.UnimplementedContentManagementServer
}

func (s *ContentManagementServer) CreateContent(ctx context.Context, input *pb.NewContent) (*pb.Content, error) {
	log.Printf("Received: %v", input.GetText())
	var content_id int32 = int32(rand.Intn(1000))

	return &pb.Content{Text: input.GetText(), Id: content_id}, nil
}

func (s *ContentManagementServer) CreateALotOfContents(stream pb.ContentManagement_CreateALotOfContentsServer) error {
	for {
		content, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Receive from client 🌊 Text: %v", content.GetText())

		var content_id int32 = int32(rand.Intn(1000))
		err = stream.Send(&pb.Content{Text: content.GetText(), Id: content_id})
		if err != nil {
			return err
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Cannot start the server %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterContentManagementServer(server, &ContentManagementServer{})

	log.Printf("Server running at %v", listen.Addr())
	err = server.Serve(listen)

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
