CREATE DATABASE ideas_db;

\c ideas_db;

-- Categories
CREATE TABLE categories (
    category_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Attributes
CREATE TABLE attributes (
    attribute_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    value TEXT
);

-- Ideas
CREATE TABLE ideas (
    idea_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category_id INT REFERENCES categories(category_id)
);

