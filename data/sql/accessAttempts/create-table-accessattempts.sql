CREATE TABLE IF NOT EXISTS accessAttempts (
    id INTEGER PRIMARY KEY,
    username VARCHAR(60),
    ip VARCHAR(60),
    attemptTimestamp TIMESTAMP,
    success BOOLEAN
);