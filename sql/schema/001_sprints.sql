-- +goose Up 
CREATE TABLE users 

CREATE TABLE sprints (
    sprint_ID INT PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    word_count INT NOT NULL
); 

-- +goose Down 
DROP TABLE sprints;