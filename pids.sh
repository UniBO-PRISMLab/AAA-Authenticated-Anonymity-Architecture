#!/bin/bash
> ~/public-identities.txt

for i in $(seq 1 100);
do
    echo "Starting PID generation $i"
    
    TEMP_DIR=$(mktemp -d)
    
    openssl genpkey -algorithm RSA -pkeyopt rsa_keygen_bits:2048 -out "$TEMP_DIR/private.key" -quiet
    openssl pkey -in "$TEMP_DIR/private.key" -pubout -out "$TEMP_DIR/public.key"

    PUBLIC_KEY_B64=$(cat "$TEMP_DIR/public.key" | base64 | tr -d '\n\r')
    
    API_RESPONSE=$(curl -X 'POST' -s \
    'http://127.0.0.1:8888/v1/identity/pid' \
    -H 'accept: application/json' \
    -H 'Content-Type: application/json' \
    -d '{
    "public_key": "'"$PUBLIC_KEY_B64"'"
    }')
    
    PID=$(echo "$API_RESPONSE" | jq -r '.pid' 2>/dev/null)
    
    if [ "$PID" == "null" ] || [ -z "$PID" ]; then
        echo "ERROR: Failed to get PID for iteration $i. Key rejected."
        ERROR_MESSAGE=$(echo "$API_RESPONSE" | jq -r '.message // .error')
        echo "API Error Message: $ERROR_MESSAGE"
        echo "$i: null (API Error)" >> ~/public-identities.txt
    else
        echo "Success. PID: $PID"
        echo "$i: $PID" >> ~/public-identities.txt
    fi
    
    rm -rf "$TEMP_DIR"
done