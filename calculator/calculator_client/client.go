package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/calculator/calculatorpb"
	"log"
)

func main() {
	fmt.Println("Hello i'm a client")

	// koneksi to server with grpc
	// Withinsecure karena selama dev masih tidak mengunakan SSL
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect %s\n", err)
	}

	/*
	Defer digunakan untuk memastikan bahwa pemanggilan fungsi
	dilakukan nanti dalam eksekusi program,	biasanya untuk tujuan cleanup.

	Defer sering digunakan seperti memastikan dan
	akhirnya akan digunakan dalam bahasa lain.
	 */
	defer cc.Close()

	c  := calculatorpb.NewCalculatorServiceClient(cc)
	//c := greetpb.NewGreetServiceClient(cc)
	//fmt.Printf("Created clien %f", c)
	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("starting to do a Unary RPC...")
	request := &calculatorpb.SumRequest{
		FirstNumber:  51,
		SecondNumber: 22,
	}
	//request := &greetpb.GreetRequest{
	//	Greeting: &greetpb.Greeting{
	//		FirstName: "Andi",
	//		LastName:  "Syafrianda",
	//	},
	//}

	//response, err := c.Greet(context.Background(), request)
	response , err := c.Sum(context.Background(), request)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v\n", err)
	}
	log.Printf("Response from Greet: %v", response.SumResult)

}
