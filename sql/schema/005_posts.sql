-- +goose Up
CREATE TABLE posts (
	id UUID PRIMARY KEY,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL,
	title TEXT NOT NULL,
	url TEXT UNIQUE NOT NULL,
	description TEXT NOT NULL,
	published_at TIMESTAMP NULL,
	feed_id UUID NOT NULL,
	CONSTRAINT fk_posts
	FOREIGN KEY (feed_id)
	REFERENCES feeds(id)
	ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;
