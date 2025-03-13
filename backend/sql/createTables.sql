CREATE TABLE IF NOT EXISTS productions (
    id BIGINT PRIMARY KEY,
    title VARCHAR(40),
    description TEXT,
    genre VARCHAR(40),
    year INTEGER
);

CREATE TABLE IF NOT EXISTS users (
    id BIGINT PRIMARY KEY,
    username VARCHAR(40) UNIQUE,
    password VARCHAR(40),
    email VARCHAR(40),
    pfp TEXT,
    description TEXT
);

CREATE TABLE IF NOT EXISTS collections (
    id BIGINT PRIMARY KEY,
    author VARCHAR(40) REFERENCES users(username),
    topic VARCHAR(40),
    message TEXT
);

CREATE TABLE IF NOT EXISTS discussions (
    id BIGINT PRIMARY KEY,
    production_id BIGINT REFERENCES productions(id),
    author VARCHAR(40) REFERENCES users(username),
    topic VARCHAR(40),
    entry_message TEXT,
    message TEXT
);

CREATE TABLE IF NOT EXISTS reviews (
    id BIGINT PRIMARY KEY,
    production_id BIGINT REFERENCES productions(id),
    author VARCHAR(40) REFERENCES users(username),
    topic VARCHAR(40),
    message TEXT
);

CREATE TABLE IF NOT EXISTS comments (
    id BIGINT PRIMARY KEY,
    type VARCHAR(20),
    text TEXT,
    entity_id BIGINT,
    author VARCHAR(40) REFERENCES users(username)
);

CREATE TABLE IF NOT EXISTS ratings (
    id BIGINT PRIMARY KEY,
    type VARCHAR(20),
    rating INTEGER,
    entity_id BIGINT,
    author VARCHAR(40) REFERENCES users(username)
);

CREATE TABLE IF NOT EXISTS collectionsProductions (
    id BIGINT PRIMARY KEY,
    collection_id BIGINT REFERENCES collections(id),
    production_id BIGINT REFERENCES productions(id),
    comment TEXT
)