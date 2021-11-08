package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AFukun/rpc-example/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:23333", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter equation: ")
	equation, _ := reader.ReadString('\n')
	client := pb.NewCalcServiceClient(conn)

	request := &pb.Request{Equation: equation}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := client.Calculate(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	result := response.GetResult()

	fmt.Printf("Result:%g\n", result)
}
