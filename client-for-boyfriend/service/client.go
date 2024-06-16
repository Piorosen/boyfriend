package service

import (
	"log"

	triton "github.com/Piorosen/boyfriend/client-for-boyfriend/grpc-client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn   *grpc.ClientConn
	client triton.GRPCInferenceServiceClient
}

func NewClient() *Client {
	return &Client{}
}

func (client *Client) Open(host string) {
	// Connect to gRPC server
	conn, err := grpc.NewClient(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Couldn't connect to endpoint %s: %v", host, err)
	}

	// Create client from gRPC server connection
	client.client = triton.NewGRPCInferenceServiceClient(conn)
	client.conn = conn
}

func (client *Client) Close() {
	client.conn.Close()
}
