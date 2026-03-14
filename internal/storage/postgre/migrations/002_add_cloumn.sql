-- +goose Up
ALTER TABLE movies ADD COLUMN kp_id INT;

-- +goose Down
ALTER TABLE movies DROP COLUMN kp_id;