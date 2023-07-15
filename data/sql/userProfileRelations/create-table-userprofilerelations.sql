CREATE TABLE IF NOT EXISTS userProfileRelations (
    id INTEGER PRIMARY KEY,
    userId INTEGER,
    profileId INTEGER,
    FOREIGN KEY (userId) REFERENCES users(id),
    FOREIGN KEY (profileId) REFERENCES profiles(id)
);