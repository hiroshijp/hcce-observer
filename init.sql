\c app

CREATE TABLE IF NOT EXISTS visitors (
    id serial PRIMARY KEY,
    mail VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS histories (
    id serial PRIMARY KEY,
    visitor_id INT REFERENCES visitors(id),
    visited_from  VARCHAR(255) NOT NULL,
    visited_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);