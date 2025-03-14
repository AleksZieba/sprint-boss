-- +goose Up 
CREATE TABLE users (
    user_name TEXT PRIMARY KEY,
    boss_fights CHAR(3) NOT NULL DEFAULT '000', 
    difficulty_level INT NOT NULL DEFAULT 2
);

CREATE TABLE sprints (
    sprint_ID INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_name TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    word_count INT NOT NULL DEFAULT 0, 
    xp INT NOT NULL DEFAULT 0,
    FOREIGN KEY (user_name) 
        REFERENCES users(user_name) 
        ON DELETE CASCADE 
        ON UPDATE CASCADE 
); 

-- +goose Down 
DROP TABLE sprints;
DROP TABLE users;
