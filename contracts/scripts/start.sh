#!/bin/sh
set -e

npx hardhat node --hostname 0.0.0.0 &
HARDHAT_PID=$!

echo "Waiting for Hardhat node to start..."
until curl -s http://127.0.0.1:8545 > /dev/null; do
  sleep 1
done

echo "Deploying contracts..."
npx hardhat run scripts/deploy.js --network localhost

wait $HARDHAT_PID
