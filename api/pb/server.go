package pb

import (
	"context"

	//pb "github.com/AliyaVV/MovieHub/api/pb"
	"github.com/AliyaVV/MovieHub/internal/model"
	"github.com/AliyaVV/MovieHub/internal/service"
)

type MovieGRPCServer struct {
	UnimplementedMovieServiceServer
	Service *service.MovieService
}

func (s *MovieGRPCServer) GetMovieById(ctx context.Context, req *GetMovieByIdRequest) (*MovieExResponse, error) {
	// movie, err := s.Service.GetMovieById(ctx, int(req.Id))
	// if err != nil {
	// 	return nil, err
	// }

	return &MovieExResponse{
		Movie: &MovieEx{
			Id:          int32(req.Id),
			Description: "Comment",
			Poster:      "test",
		},
	}, nil
}

func (s *MovieGRPCServer) GetMovies(ctx context.Context, req *Empty) (*MovieShortListResponse, error) {

	// movies, err := s.Service.GetMovies(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	return &MovieShortListResponse{
		Movies: []*MovieShort{
			{
				Id:        1,
				Runame:    "Test Movie",
				Enname:    "Test Movie",
				Movietype: "movie",
				Movieyear: 2024,
			},
		},
	}, nil
}

// --- CreateMovie ---
func (s *MovieGRPCServer) CreateMovie(ctx context.Context, req *CreateMovieRequest) (*MovieShortResponse, error) {

	m := req.Movie

	movie := &model.Movie_ex{
		Movie_short: model.Movie_short{
			Id:        int(m.Id),
			Runame:    m.Runame,
			EnName:    m.Enname,
			MovieType: m.Movietype,
			MovieYear: int(m.Movieyear),
		},
	}

	id, err := s.Service.MovieRepo.SaveMovie(ctx, movie)
	if err != nil {
		return nil, err
	}

	return &MovieShortResponse{
		Movie: &MovieShort{
			Id: int32(id),
		},
	}, nil
}
