package main

import (
	"context"
	"log"
	"net"
	"time"
	"errors"
	"google.golang.org/grpc"
	pb "grpcserver.com/pkg" // Import your generated protobuf package
)

// Define a struct to implement the Greeter service
type greeterServer struct{
    pb.UnimplementedGreeterServer
}
// Define a struct to implement the Calculator service
type calculatorServer struct{
    pb.UnimplementedCalculatorServer
}
// Implement the SayHello RPC method
func (s *greeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    // return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
	if in.GetName() == "shubham" {
        // Create a new context with a timeout of 20 seconds
        log.Printf("Received: %v", in.GetName())
        timeoutCtx, cancel := context.WithTimeout(ctx, 20*time.Second)
        defer cancel()

        // Create a channel to signal when the response is ready
        ch := make(chan *pb.HelloReply)

        // Run the server logic in a goroutine
        go func() {
            // Simulate a delay of 20 seconds for requests with the name "shubham"
            time.Sleep(20 * time.Second)

            // Send the response back on the channel
            ch <- &pb.HelloReply{Message: "Hello " + in.GetName()}
        }()

        // Wait for either the response or the context timeout
        select {
        case <-timeoutCtx.Done():
            return nil, errors.New("request timed out")
        case reply := <-ch:
            return reply, nil
        }
    }

    // For requests with other names, proceed normally
    log.Printf("Received: %v", in.GetName())
    return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// Implement the Add RPC method
func (s *calculatorServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	// Perform addition
	result := req.GetNum1() + req.GetNum2()

	// Construct the response message
	resp := &pb.AddResponse{
		Result: result,
	}
	return resp, nil
}

func main() {
	// Create a listener on TCP port 50053
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a gRPC server
	s := grpc.NewServer()

	// Register the Greeter service with the server
	pb.RegisterGreeterServer(s, &greeterServer{})

	// Register the Calculator service with the server
	pb.RegisterCalculatorServer(s, &calculatorServer{})

	// Serve the gRPC server
	log.Println("Server is running on port 50053...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

