package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc_patterns/client/ecommerce"
	"log"
	"time"
)

const(
	address="localhost:50051"
)

func main(){
	//setup connection with the server
	conn, err:=grpc.Dial(address, grpc.WithInsecure())
	if err!=nil{
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//create a stub, that contains all the remote methods to invoke the server
	c:=pb.NewProductInfoClient(conn)

	name:="Apple iPhone 11"
	description:="Meet Apple iPhone 11. All new dual camera system...."

	//create a context to pass with the remote call
	ctx, cancel:=context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err:=c.AddProduct(ctx, &pb.Product{Name: name, Description: description})
	if err!=nil{
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf("Product ID: %s added successfully", r.Value)

	product, err:=c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err!=nil{
		log.Fatalf("Could not get product: %v", err)
	}
	log.Printf("Product: %s", product.String())
}