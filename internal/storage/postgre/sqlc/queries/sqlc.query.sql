-- name: SaveMovie :one
INSERT INTO movies (
    kp_id,
    ru_name,
    en_name,
    movie_type,
    year,
    description,
    top250,
    budget,
    revenue,
    tmdb_id
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
RETURNING id;

-- name: SaveRating :exec
INSERT INTO ratings(
    movie_id,
    kp,
    tmdb,
    film_critic,
    russian_film_critics
) VALUES ($1,$2,$3,$4,$5);

-- name: SaveGenre :exec
INSERT INTO movie_genres (
    movie_id,
    genre_id
)
VALUES ($1,$2);

-- name: SaveCast :exec
INSERT INTO casts(
    movie_id,
    name,
    en_name,
    profession,
    description
) VALUES ($1,$2,$3,$4,$5);


-- name: GetMovie :one
SELECT
    mv.kp_id,
    mv.id,
    mv.ru_name,
    mv.en_name,
    mv.movie_type,
    mv.description,
    mv.top250,
    mv.budget,
    mv.revenue,
    mv.tmdb_id,
    mv.year,
    rs.kp,
    rs.tmdb,
    rs.film_critic,
    rs.russian_film_critics
FROM movies mv
LEFT JOIN ratings rs
ON mv.id = rs.movie_id
WHERE mv.kp_id = $1;

-- name: GetGenres :many
SELECT
    g.id,
    g.name
FROM movie_genres mg
JOIN genres g
ON g.id = mg.genre_id
WHERE mg.movie_id = $1;

-- name: GetCast :many
SELECT
    id,
    movie_id,
    name,
    en_name,
    profession,
    description
FROM casts
WHERE movie_id = $1;

-- name: GetListMovies :many
SELECT
    mv.id,
    mv.kp_id,
    mv.ru_name,
    mv.en_name,
    mv.movie_type,
    mv.tmdb_id,
    mv.year,
    rs.kp,
    rs.tmdb
FROM movies mv
LEFT JOIN ratings rs
ON mv.id = rs.movie_id
;

-- name: GetGenreByName :one
SELECT id FROM genres
WHERE name = $1;

-- name: SaveAwards :exec
INSERT INTO awards(
    movie_id,
    award
) VALUES ($1,$2);


