-- name: StartSprint :exec
INSERT INTO users (user_ID, user_name, server_name)
VALUES ($1 || '\\' || $2, $1, $2)
ON CONFLICT (user_ID) DO UPDATE
    SET user_name = EXCLUDED.user_name;

INSERT INTO sprints (user_name, created_at, updated_at) 
VALUES ($1 || '\\' || $2, now(), now());
