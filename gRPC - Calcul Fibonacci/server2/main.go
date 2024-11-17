package main

import (
	"context"
	pb "fibonacci/fibonacci"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFibonacciServiceServer
}

func (s *server) GetFibonacci(ctx context.Context, req *pb.FibonacciRequest) (*pb.FibonacciResponse, error) {
	fmt.Println(ctx)
	n := req.Number
	res := fibonacci(n)
	return &pb.FibonacciResponse{Sequence: res}, nil
}

func fibonacci(n int32) []int32 {
	seq := []int32{0, 1}
	for i := 2; i <= int(n); i++ {
		seq = append(seq, seq[i-1]+seq[i-2])
	}
	return seq[:n]
}

func main() {
	lis, err := net.Listen("tcp", ":50123")
	if err != nil {
		fmt.Println("Can't connect to port 50123")
		return
	}

	s := grpc.NewServer()
	pb.RegisterFibonacciServiceServer(s, &server{})
	fmt.Println("Connected")
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
