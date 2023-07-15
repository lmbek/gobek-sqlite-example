CREATE TABLE IF NOT EXISTS sessions (
    id INTEGER PRIMARY KEY,
    name VARCHAR(60),
    ip VARCHAR(60),
    token TEXT,
    startTime TIMESTAMP,
    endTime TIMESTAMP
);