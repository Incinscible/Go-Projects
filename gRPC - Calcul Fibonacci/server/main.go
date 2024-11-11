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
	n := req.Number
	sequence := fibonacci(n)
	return &pb.FibonacciResponse{Sequence: sequence}, nil
}

func fibonacci(n int32) []int32 {
	seq := []int32{0, 1}
	for i := 2; i <= int(n); i++ {
		seq = append(seq, seq[i-1]+seq[i-2])
	}
	return seq[:n]
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterFibonacciServiceServer(s, &server{})
	fmt.Print("Serveur gRPC en écoute sur le port 50051")
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
