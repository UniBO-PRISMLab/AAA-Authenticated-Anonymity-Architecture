INSERT INTO sac (sac, expiration, sid)
VALUES ($1, $2, $3)
RETURNING sac, expiration, sid;