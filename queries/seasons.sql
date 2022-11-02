-- name: GetSeasonsForShow :many
SELECT * FROM seasons WHERE show_id = ?;