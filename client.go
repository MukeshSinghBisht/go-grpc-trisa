package main

import (
	"log"

	"google.golang.org/grpc"

	pb "path/to/your/proto/package"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create client instances
	trisaNetworkClient := pb.NewTrisaNetworkClient(conn)
	trisaHealthClient := pb.NewTRISAHealthClient(conn)

	// Use the clients to call the gRPC methods
	// Example:
	// trisaNetworkClient.Transfer(context.Background(), &pb.SecureEnvelope{})
	// trisaHealthClient.Status(context.Background(), &pb.HealthCheck{})
}
