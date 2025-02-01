package main

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	desc "github.com/UraharaKiska/chat-server/pkg/chat_v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const grpcPort = 50054

type server struct {
	desc.UnimplementedChatV1Server
}

func (s *server) Create(ctx context.Context, in *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Created: %v", in.GetUsernames())

	return &desc.CreateResponse{
		Id: int64(gofakeit.Number(1000, 10000)),
	}, nil
}

func (s *server) Delete(ctx context.Context, in *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Deleted: %v", in.GetId())

	return &emptypb.Empty{}, nil
}

func (s *server) SendMessage(ctx context.Context, in *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Send message to %+v", in)

	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("grpc server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve: ", err)
	}
}
