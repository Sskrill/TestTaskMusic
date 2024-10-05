-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS song (
    id SERIAL PRIMARY KEY,
    song_name VARCHAR(255) NOT NULL ,
    performer_name VARCHAR(255) NOT NULL ,
    link VARCHAR(255),
    song_text TEXT,
    release_date DATE,
    created_at TIMESTAMP NOT NULL
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS song;
-- +goose StatementEnd
