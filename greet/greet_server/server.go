package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/greet/greetpb"
	"log"
	"net"
)

type server struct {
}

// unary API Server implementation
func (s *server) Greet(_ context.Context, request *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("fungsi Greet telah di invoke with %v\n", request)
	firstName := request.GetGreeting().GetFirstName()
	result := "Hello "+ firstName

	response := &greetpb.GreetResponse{
		Result: result,
	}
	return response, nil
}

func main() {
	fmt.Println("Hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listend : %v\n", err)
	}

	// inisiasi grpc
	s := grpc.NewServer()
	// RegisterGreetServiceServer membuat &server{}
	// otomatis implement method struct & Greet
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}

}
