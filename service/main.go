package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "grpc_patterns/service/ecommerce"
	"log"
	"net"
)

const (
	port=":50051"
)

func main(){
	lis, err:=net.Listen("tcp", port)
	if err!=nil{
		log.Fatalf("Failed to listen: %v", err )
	}
	s:=grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})
	reflection.Register(s)
	log.Printf("Starting gRPC listener on port "+port)

	if err:=s.Serve(lis); err!=nil{
		log.Fatalf("Failed to serve: %v", err)
	}
}

