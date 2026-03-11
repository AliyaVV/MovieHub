-- +goose Up
CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    ru_name TEXT,
    en_name TEXT,
    movie_type TEXT,
    year INT,
    description TEXT,
    top250 INT,
    budget INT,
    revenue INT,
    tmdb_id INT
);

CREATE TABLE ratings (
    id SERIAL PRIMARY KEY,
    movie_id INT REFERENCES movies(id),
    kp FLOAT,
    tmdb FLOAT,
    film_critic FLOAT,
    russian_film_critics FLOAT
);
CREATE TABLE genres (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE
);
INSERT INTO genres(name)
VALUES
    ('аниме'),
    ('драма'),
    ('мелодрама'),
    ('боевик'),
    ('вестерн'),
    ('военный'),
    ('детектив'),
    ('детский'),
    ('документальный'),
    ('исторический'),
    ('комедия'),
    ('мультфильм'),
    ('мюзикл'),
    ('фэнтези'),
    ('приключения'),
    ('фантастика'),
    ('биография'),
    ('для взрослых'),
    ('история'),
    ('короткометражка'),
    ('криминал'),
    ('музыка'),
    ('новости'),
    ('семейный'),
    ('спорт'),
    ('ток-шоу'),
    ('триллер'),
    ('ужасы'),
    ('фильм-нуар'),
    ('церемония');


CREATE TABLE movie_genres (
    movie_id INT REFERENCES movies(id),
    genre_id INT REFERENCES genres(id),
    PRIMARY KEY(movie_id, genre_id)
);
CREATE TABLE casts (
    id SERIAL PRIMARY KEY,
    movie_id INT REFERENCES movies(id),
    name TEXT,
    en_name TEXT,
    profession TEXT,
    description TEXT
);
CREATE INDEX movie_id_idx ON ratings(movie_id);
CREATE INDEX movie_name_idx ON movies(ru_name);
CREATE INDEX cast_movie_id_idx ON casts(movie_id);

-- +goose Down
DROP TABLE IF EXISTS movie_genres;
DROP TABLE IF EXISTS casts;
DROP TABLE IF EXISTS ratings;
DROP TABLE IF EXISTS genres;
DROP TABLE IF EXISTS movies;