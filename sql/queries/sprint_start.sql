-- name: StartSprint :exec
BEGIN;
INSERT INTO users (user_name)
VALUES ($1)
ON CONFLICT (user_name) DO NOTHING;

INSERT INTO sprints (user_name, created_at, updated_at) 
VALUES ($1, now(), now());
COMMIT;