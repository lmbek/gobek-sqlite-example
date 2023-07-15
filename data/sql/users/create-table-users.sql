CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    username VARCHAR(60),
    password VARCHAR(255),
    salt VARCHAR(60),
    email VARCHAR(60),
    phone VARCHAR(60)
);