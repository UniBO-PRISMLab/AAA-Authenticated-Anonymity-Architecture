CREATE TABLE IF NOT EXISTS pid (
    id SERIAL PRIMARY KEY,
    pid TEXT NOT NULL,
    public_key TEXT NOT NULL,
    nonce TEXT NOT NULL,
    timestamp timestamp default current_timestamp
);
