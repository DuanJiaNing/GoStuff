package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "GoStuff/grpc/services"
)

func main() {
	conn, err := grpc.Dial("localhost:3332", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAccountClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AddAccount(ctx, &pb.AccountRequest{
		Name:    "Jim",
		Id:      &pb.Id{Id: int32(999)},
		Married: true})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("After add account: %s", r)

	r, err = c.GetAccount(ctx, &pb.Id{Id: int32(999)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("After get account: %s", r)
}
