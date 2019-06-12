package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"google.golang.org/grpc"

	pb "GoStuff/grpc/services"
)

type tokenAuth struct {
	token string
}

// Return value is mapped to request headers.
func (t tokenAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	return map[string]string{
		"authorization": "Bearer " + t.token,
	}, nil
}

func (tokenAuth) RequireTransportSecurity() bool {
	return true
}

func main() {
	conn, err := grpc.Dial("localhost:3332", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	send(context.Background(), conn)
}

func send(ctx context.Context, conn *grpc.ClientConn) {
	c := pb.NewAccountClient(conn)

	ctx, cancel := context.WithTimeout(ctx, time.Second)
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

func handle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	token := "test"

	conn, err := grpc.DialContext(ctx, net.JoinHostPort("localhost", "3332"),
		//grpc.WithTransportCredentials(credentials.NewClientTLSFromCert(insecure.CertPool, "")),
		grpc.WithPerRPCCredentials(tokenAuth{
			token: token,
		}),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	send(ctx, conn)

}
