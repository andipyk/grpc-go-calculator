package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/calculator/calculatorpb"
	"log"
	"net"
)

type server struct {}

func (s server) Sum(ctx context.Context, request *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Received sum RPC: %v\n", request)
	firstNumber := request.FirstNumber
	secondNumber := request.SecondNumber
	sum := firstNumber + secondNumber
	response := &calculatorpb.SumResponse{
		SumResult: sum,
	}
	return response, nil
}

// unary API Server implementation
func main() {
	fmt.Println("Hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listend : %v\n", err)
	}

	// inisiasi grpc
	s := grpc.NewServer()
	// RegisterGreetServiceServer membuat &server{}
	// otomatis implement method struct & Calculator
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v\n", err)
	}

}
