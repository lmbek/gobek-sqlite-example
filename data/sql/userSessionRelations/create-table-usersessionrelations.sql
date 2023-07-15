CREATE TABLE IF NOT EXISTS userSessionRelations (
    id INTEGER PRIMARY KEY,
    userId INTEGER,
    sessionId INTEGER,
    FOREIGN KEY (userId) REFERENCES users(id),
    FOREIGN KEY (sessionId) REFERENCES sessions(id)
);