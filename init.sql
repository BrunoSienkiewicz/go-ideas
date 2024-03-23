-- Create database
CREATE DATABASE ideas_db;

\c ideas_db;

CREATE SCHEMA data;

-- Categories
CREATE TABLE data.categories (
    category_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Attributes
CREATE TABLE data.attributes (
    attribute_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    value TEXT
);

-- Ideas
CREATE TABLE data.ideas (
    idea_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category_id INT REFERENCES data.categories(category_id)
);

