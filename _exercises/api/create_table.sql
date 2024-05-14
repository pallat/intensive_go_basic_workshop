CREATE TABLE languages (
    id INTEGER primary key autoincrement,
    name TEXT,
    imageUrl TEXT,
    created_at TEXT, -- ISO8601: YYYY-MM-DD HH:MM:SS.SSS
    updated_at TEXT,
    deleted_at TEXT
);
