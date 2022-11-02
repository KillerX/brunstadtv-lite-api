-- name: ListShows :many
SELECT * FROM shows;

-- name: GetShowByID :one
SELECT * FROM shows WHERE id = ?;
