package main

import (
	"context"
	"log"
	"time"

	pb "github.com/AliyaVV/MovieHub/api/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal("cannot connect:", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Println("Error close redis", err)
		}
	}()

	client := pb.NewMovieServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := client.GetMovieById(ctx, &pb.GetMovieByIdRequest{
		Id: 101,
	})
	if err != nil {
		log.Fatal("error:", err)
	}

	log.Println("Movie:", resp.Movie)

	list, err := client.GetMovies(ctx, &pb.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Movies:", list.Movies)
}
