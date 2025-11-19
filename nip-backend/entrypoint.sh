#!/bin/sh
set -e

INSTANCE_ID="${TASK_SLOT:-1}"

echo "I am starting as Instance ID: $INSTANCE_ID"

CONFIG_FILE="/app/configs/.env.instance${INSTANCE_ID}"

if [ -f "$CONFIG_FILE" ]; then
    echo "Loading config: $CONFIG_FILE"
    set -a
    . "$CONFIG_FILE"
    set +a
else
    echo "ERROR: Config file $CONFIG_FILE not found!"
    exit 1
fi

exec /app/nip-backend