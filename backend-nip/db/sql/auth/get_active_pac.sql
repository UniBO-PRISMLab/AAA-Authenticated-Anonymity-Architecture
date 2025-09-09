SELECT pac, expiration
FROM active_pac
WHERE pid = $1 AND pac = $2;