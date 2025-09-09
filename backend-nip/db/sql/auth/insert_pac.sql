INSERT INTO pac (pac, expiration, pid)
VALUES ($1, $2, $3)
RETURNING pac, expiration, pid;