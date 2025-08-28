INSERT INTO pid (pid, public_key)
VALUES ($1, $2)
RETURNING pid, public_key;