CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS pid (
    pid TEXT PRIMARY KEY,
    public_key TEXT NOT NULL,
    public_key_hash BYTEA GENERATED ALWAYS AS (digest(public_key, 'sha256')) STORED UNIQUE,
    nonce TEXT NOT NULL,
    timestamp TIMESTAMP DEFAULT current_timestamp
);

CREATE UNIQUE INDEX idx_pid_public_key_hash ON pid (public_key_hash);

CREATE TABLE IF NOT EXISTS pac (
    id SERIAL PRIMARY KEY,
    pac INTEGER NOT NULL,
    expiration TIMESTAMP NOT NULL,
    pid TEXT NOT NULL REFERENCES pid(pid),
    timestamp TIMESTAMP DEFAULT current_timestamp
);

CREATE VIEW active_pac AS
SELECT *
FROM pac
WHERE expiration > now();

CREATE INDEX idx_pac_pid ON pac (pid);