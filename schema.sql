CREATE DATABASE "GoNews"

CREATE TABLE authors (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    author_id INTEGER REFERENCES authors(id) NOT NULL,
    title TEXT  NOT NULL,
    content TEXT NOT NULL,
    created_at BIGINT NOT NULL
);