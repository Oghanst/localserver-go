package main

import (
	"context"
	"log"
	"time"

	// "net"

	pb "github.com/Oghanst/localserver-go.git/proto/securityproto"

	// "github.com/cloudevents/sdk-go/v2/event"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:7202"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSecurityPullServiceClient(conn)

	// Contact the server and print out its response.

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	

	r, err := c.PullConfig(ctx, &pb.SecurityStateRequest{
		State: &pb.SecurityState{
			Timestamp: time.Now().Unix(),
			Name: "test",
			Available: 0,
		},
		Ip: "1.2.4.6",
		Hostname: "localhost",
	})
	if err != nil {
		log.Fatalf("could not get response: %v", err)
	}
	log.Printf("Response: %s", r)
}
