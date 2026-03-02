package mongo

import (
	"context"
	"errors"
	"fmt"

	"github.com/AliyaVV/MovieHub/internal/model"
	"github.com/AliyaVV/MovieHub/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoMovieRepository struct {
	collection *mongo.Collection
}

func NewMongoMovieRepository(db *mongo.Database) repository.MovieRepository {
	return &MongoMovieRepository{
		collection: db.Collection("movies"),
	}
}

func (r *MongoMovieRepository) Upsert(ctx context.Context, movie *model.Movie_short) error {
	filter := bson.M{"id": movie.Id}
	update := bson.M{
		"$set": movie,
	}

	opts := options.Update().SetUpsert(true)
	result, err := r.collection.UpdateOne(ctx, filter, update, opts)
	fmt.Println("Matched:", result.MatchedCount)
	fmt.Println("Upserted:", result.UpsertedCount)

	return err
}

func (r *MongoMovieRepository) GetById(ctx context.Context, id int) (*model.Movie_short, error) {
	filter := bson.M{"id": id}

	var movie model.Movie_short

	err := r.collection.FindOne(ctx, filter).Decode(&movie)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}

	return &movie, nil
}

func InitMongo() (*mongo.Client, error) {
	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI("mongodb://localhost:27017"),
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}
