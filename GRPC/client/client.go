package main

import (
	"context"
	"log"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc"
	pb "veltrisclient.com/pkg"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client instance
	client := pb.NewGreeterClient(conn)
	calclient := pb.NewCalculatorClient(conn)

	// Create a context with a timeout of 30 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create a channel to receive responses
	ch := make(chan *pb.HelloReply)

	// Wait group to wait for all server responses to be received
	var wg sync.WaitGroup

	// Concurrently make requests with different names
	for i := 0; i < 9; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			// Make a request with a different name
			str := strconv.Itoa(i)
			resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "other" + str})
			if err != nil {
				log.Printf("RPC failed for other%s: %v", str, err)
			} else {
				ch <- resp
			}
		}(i)
	}

	// Concurrently make request with name "shubham"
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Create a context with a timeout of 20 seconds for the request with the name "shubham"
		ctxShubham, cancelShubham := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancelShubham()

		// Make the RPC call with the new context
		log.Println("sent req with name shubham hold for 20sec waiting to come resp from server")
		resp, err := client.SayHello(ctxShubham, &pb.HelloRequest{Name: "shubham"})
		if err != nil {
			log.Printf("RPC failed for shubham: %v", err)
		} else {
			ch <- resp
		}
	}()

	// Call the add function using a goroutine
	resultCh := make(chan int32)
	go func() {
		result, err := add(ctx, calclient)
		if err != nil {
			log.Fatalf("Error while calling Add RPC: %v", err)
		}
		resultCh <- result
	}()

	// Start a goroutine to wait for all server responses to be received
	go func() {
		wg.Wait()
		close(ch) // Close the channel to signal that all responses have been received
	}()

	// Wait for responses from all requests
	for {
		select {
		case resp, ok := <-ch:
			if !ok {
				// Channel closed, all responses received
				return
			}
			log.Printf("Response from server: %s", resp.GetMessage())
		case result := <-resultCh:
			log.Printf("Result of addition: %d", result)
		case <-ctx.Done():
			log.Println("Context timeout exceeded")
			return
		}
	}
}

func add(ctx context.Context, client pb.CalculatorClient) (int32, error) {
	// Define the request
	req := &pb.AddRequest{
		Num1: 10, // Example value for Num1
		Num2: 20, // Example value for Num2
	}

	// Call the Add method on the server
	resp, err := client.Add(ctx, req)
	if err != nil {
		return 0, err
	}

	return resp.Result, nil
}

