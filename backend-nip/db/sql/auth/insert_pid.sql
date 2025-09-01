INSERT INTO pid (pid, public_key, nonce)
VALUES ($1, $2, $3)
RETURNING pid, public_key, nonce;