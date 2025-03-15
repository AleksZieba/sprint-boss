-- +goose Up 
CREATE TABLE users (
    user_ID TEXT PRIMARY KEY,
    user_name TEXT NOT NULL,
    server_name TEXT NOT NULL,
    difficulty_level INT NOT NULL DEFAULT 2
);

CREATE TABLE sprints (
    sprint_ID INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_ID TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    word_count INT NOT NULL DEFAULT 0, 
    xp INT NOT NULL DEFAULT 0,
    FOREIGN KEY (user_ID) 
        REFERENCES users(user_ID) 
        ON DELETE CASCADE 
        ON UPDATE CASCADE 
); 

-- +goose Down 
DROP TABLE sprints;
DROP TABLE users;
