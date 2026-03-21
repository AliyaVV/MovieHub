package main

import (
	"log"
	"net"

	pb "github.com/AliyaVV/MovieHub/api/pb"
	"github.com/AliyaVV/MovieHub/internal/service"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	movieService := &service.MovieService{}
	pb.RegisterMovieServiceServer(grpcServer, &pb.MovieGRPCServer{
		Service: movieService,
	})

	log.Println("Server started on port 50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
