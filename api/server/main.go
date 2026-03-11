package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	//pb.RegisterMovieServiceServer(grpcServer, &pb.MovieGRPCServer{})

	log.Println("Server started on port 50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
