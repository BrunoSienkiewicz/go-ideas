-- Create schema
CREATE SCHEMA IF NOT EXISTS ideas;

-- Categories
CREATE TABLE ideas.categories (
    category_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

-- Ideas
CREATE TABLE ideas.ideas (
    idea_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category_id INT REFERENCES ideas.categories(category_id)
);

-- Attributes
CREATE TABLE ideas.attributes (
    attribute_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    value TEXT,
    idea_id INT REFERENCES ideas.ideas(idea_id)
);



