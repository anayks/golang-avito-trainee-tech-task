CREATE TABLE users (
	id BIGSERIAL NOT NULL PRIMARY KEY, 
	username VARCHAR NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);