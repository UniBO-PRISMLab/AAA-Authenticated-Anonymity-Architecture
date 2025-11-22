#!/bin/sh
N=$1

echo "Generating configuration files for $1 instances..."

ACCOUNTS_FILE="./accounts.txt"

addresses=$(awk '/^Private Keys/{exit} /^\([0-9]+\) 0x/ {print $2}' "$ACCOUNTS_FILE")
privkeys=$(awk '/^Private Keys/{flag=1; next} flag && /^\([0-9]+\) 0x/ {print $2}' "$ACCOUNTS_FILE")
addr_array=($addresses)
pk_array=($privkeys)

for i in $(seq 1 $N);
do
    echo "Generating file .env.instance$i"
    openssl genpkey -algorithm RSA -pkeyopt rsa_keygen_bits:2048 -out private.key -quiet
    openssl pkey -in private.key -pubout -out public.key

    PRIVATE_KEY_B64=$(base64 -i private.key | tr -d '\n')
    PUBLIC_KEY_B64=$(base64 -i public.key | tr -d '\n')
    SK=$(openssl rand -base64 48)

    BLOCKCHAIN_ADDRESS="${addr_array[$((i-1))]}"
    BLOCKCHAIN_PRIVATE_KEY="${pk_array[$((i-1))]}"

    DB_NAME=$(cat db_name.txt)
    DB_PASSWORD=$(cat db_password.txt)
    DB_USER=$(cat db_user.txt)

    if [ -z "$DB_NAME" ] || [ -z "$DB_PASSWORD" ] || [ -z "$DB_USER" ]; then
        echo "Error: Database credentials are not set properly."
        exit 1
    fi

    cat > ./nip-backend/configs/.env.instance$i <<EOF
GIN_MODE="release"
DATABASE_URL="postgres://$DB_USER:$DB_PASSWORD@db:5432/$DB_NAME"
SK="$SK"
ETH_NODE_URL="ws://contracts:8545"
CONTRACT_ADDRESS="0x5FbDB2315678afecb367f032d93F642f64180aa3"
PUBLIC_KEY="$PUBLIC_KEY_B64"
PRIVATE_KEY="$PRIVATE_KEY_B64"
BLOCKCHAIN_PRIVATE_KEY="$BLOCKCHAIN_PRIVATE_KEY"
BLOCKCHAIN_ADDRESS="$BLOCKCHAIN_ADDRESS"
HTTP_HOST="0.0.0.0"
HTTP_PORT="8888"
EOF

done

rm -f private.key public.key
