#!/bin/sh
set -e

anvil --host 0.0.0.0 --chain-id 31337 -a 36 --gas-limit 100000000000 &
ANVIL_PID=$!

echo "Waiting for Anvil node to start..."
until curl -s http://127.0.0.1:8545 > /dev/null; do
  sleep 1
done

echo "Deploying contracts using Hardhat..."
npx hardhat run scripts/deploy.js --network localhost

echo "Contracts deployed!"
wait $ANVIL_PID
