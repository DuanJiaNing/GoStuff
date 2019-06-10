package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "GoStuff/grpc/services"
)

type server struct{}

var (
	accounts map[int32]*pb.AccountRequest
)

func init() {
	accounts = make(map[int32]*pb.AccountRequest)
}

func (s *server) GetAccount(ctx context.Context, id *pb.Id) (*pb.AccountReply, error) {
	log.Printf("GetAccount: %v", id)
	a := accounts[id.Id]

	return &pb.AccountReply{
		Id:      a.Id,
		Name:    a.Name,
		Married: a.Married}, nil
}

func (s *server) AddAccount(ctx context.Context, ac *pb.AccountRequest) (*pb.AccountReply, error) {
	accounts[ac.Id.Id] = ac

	return &pb.AccountReply{
		Id:         ac.Id,
		Name:       ac.Name,
		Married:    ac.Married,
		CreateTime: int64(time.Now().Nanosecond())}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":3332")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAccountServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
