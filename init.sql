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

INSERT INTO visitors (mail) values ('foo@example.com');
INSERT INTO visitors (mail) values ('bar@example.com');

INSERT INTO histories (visitor_id, visited_from) values (1, 'http://example.com');
INSERT INTO histories (visitor_id, visited_from) values (2, 'http://example.com');