package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/AFukun/rpc-example/pb"
	"google.golang.org/grpc"
)

func replicateCalc(equation string, portList []string) (float32, int) {
	results := make(chan float32)
	request := &pb.Request{Equation: equation}

	var wg sync.WaitGroup
	for _, port := range portList {
		wg.Add(1)
		go func(port string) {
			defer wg.Done()
			conn, err := grpc.Dial("127.0.0.1:"+port, grpc.WithInsecure())
			if err != nil {
				return
			}
			client := pb.NewCalcServiceClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			response, err := client.Calculate(ctx, request)
			if err != nil {
				return
			}
			result := response.GetResult()
			results <- result
		}(port)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	counter := make(map[float32]int)
	responseCount := 0
	for result := range results {
		counter[result]++
		responseCount++
	}
	var bestResult float32
	max := 0
	for result, count := range counter {
		if count > max {
			max = count
			bestResult = result
		}
	}
	return bestResult, responseCount
}

func main() {
	portList := []string{"9000", "9001", "9002", "9003", "9004", "9005"}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter equation: ")
	equation, _ := reader.ReadString('\n')

	result, responseCount := replicateCalc(equation, portList)

	if responseCount != 0 {
		fmt.Printf("Total %d Server Response, most with Result:%g\n", responseCount, result)

	} else {
		fmt.Println("Error: no server online")
	}
}
