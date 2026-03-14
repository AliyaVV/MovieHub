package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoMovieRepository struct {
	collection *mongo.Collection
}

/*
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

func (r *MongoMovieRepository) GetAll(ctx context.Context) ([]*model.Movie_short, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var movies []*model.Movie_short
	for cursor.Next(ctx) {
		var movie model.Movie_short
		if err := cursor.Decode(&movie); err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}

	return movies, nil
}

func (r *MongoMovieRepository) Create(ctx context.Context, movie *model.Movie_short) error {
	_, err := r.collection.InsertOne(ctx, movie)
	return err
}
*/
