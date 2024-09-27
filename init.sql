\c app

CREATE TABLE IF NOT EXISTS visitors (
    id serial PRIMARY KEY,
    mail VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS histories (
    id serial PRIMARY KEY,
    visitor_id INT REFERENCES visitors(id),
    visited_from  VARCHAR(255) NOT NULL,
    visited_at TIMESTAMP  
);

CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    is_admin BOOLEAN NOT NULL
);

--- Insert some data
INSERT INTO visitors (mail) values ('test@example.com');
INSERT INTO histories (visitor_id, visited_from, visited_at) values (1,'http://example.com', NOW());
INSERT INTO users (name, password, is_admin) values ('hiroshijp', 'password', true);