
// package pb

// import (
// 	"context"

// 	pb "github.com/AliyaVV/MovieHub/api"
// 	"github.com/AliyaVV/MovieHub/internal/service"
// )

// type MovieGRPCServer struct {
// 	pb.UnimplementedMovieServiceServer
// 	service *service.MovieService
// }

// func (s *MovieGRPCServer) GetMovieById(ctx context.Context, req *pb.GetMovieByIdRequest) (*pb.MovieExResponse, error) {
// 	movie, err := s.service.GetMovieById(ctx, int(req.Id))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &pb.MovieExResponse{
// 		Movie: &pb.MovieEx{
// 			Id:          int32(movie.Id),
// 			Description: movie.Description,
// 			Poster:      movie.Poster,
// 		},
// 	}, nil
// }

// func (s *MovieGRPCServer) GetMovies(ctx context.Context, req *pb.Empty) (*pb.MovieShortListResponse, error) {

// 	movies, err := s.service.GetMovies(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var pbMovies []*pb.MovieShort

// 	for _, m := range movies {
// 		pbMovies = append(pbMovies, &pb.MovieShort{
// 			Id: int32(m.Id),
// 		})
// 	}

// 	return &pb.MovieShortListResponse{
// 		Movies: pbMovies,
// 	}, nil
// }
