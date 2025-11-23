#!/bin/bash
for i in $(seq 1 1000);
do
    echo "Generating pid $i"
    openssl genpkey -algorithm RSA -pkeyopt rsa_keygen_bits:2048 -out private.key -quiet
    openssl pkey -in private.key -pubout -out public.key

    PRIVATE_KEY_B64=$(base64 -i private.key | tr -d '\n')
    PUBLIC_KEY_B64=$(base64 -i public.key | tr -d '\n')

    PID=$(curl -X 'POST' \
    'http://127.0.0.1:8888/v1/identity/pid' \
    -H 'accept: application/json' \
    -H 'Content-Type: application/json' \
    -d '{
    "public_key": "'"$PUBLIC_KEY_B64"'"
    }' | jq -r '.pid' )

    echo "$i: $PID" >> ~/public-identities.txt
done

rm -f private.key public.key

