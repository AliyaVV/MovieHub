-- +goose Up
CREATE TABLE awards (
    id SERIAL PRIMARY KEY,
    movie_id INT REFERENCES movies(id),
    award TEXT
);

-- +goose Down
DROP TABLE IF EXISTS awards;