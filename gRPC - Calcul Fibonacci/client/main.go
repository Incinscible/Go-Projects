package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	pb "fibonacci/fibonacci"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50123", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Erreur lors de la connexion : %v", err)
	}
	defer conn.Close()

	client := pb.NewFibonacciServiceClient(conn)

	number, _ := strconv.Atoi(os.Args[1])
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.GetFibonacci(ctx, &pb.FibonacciRequest{Number: int32(number)})
	if err != nil {
		log.Fatalf("Erreur lors de la requête : %v", err)
	}

	fmt.Printf("Suite de Fibonacci pour n=%d : %v\n", number, response.Sequence)
}
