package main

import (
	"context"
	"log"
	"net"

	"github.com/AFukun/rpc-example/pb"
	"github.com/AFukun/rpc-example/server/calculator"
	"google.golang.org/grpc"
)

// server implements EchoServer.
type server struct{}

func (s *server) Calculate(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	result, err := calculator.Calc(request.GetEquation())
	if err != nil {
		return &pb.Response{Result: 0.0}, err
	}
	return &pb.Response{Result: float32(result)}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":23333")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterCalcServiceServer(s, &server{})

	if e := s.Serve(listener); e != nil {
		log.Fatal(e)
	}
}
