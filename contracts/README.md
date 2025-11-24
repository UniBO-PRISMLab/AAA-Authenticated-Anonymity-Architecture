# Contracts

[![Push contracts to ghcr.io](https://github.com/UniBO-PRISMLab/AAA-Authenticated-Anonymity-Architecture/actions/workflows/contracts-ghcr.yaml/badge.svg)](https://github.com/UniBO-PRISMLab/AAA-Authenticated-Anonymity-Architecture/actions/workflows/contracts-ghcr.yaml)
[![Contracts Tests](https://github.com/UniBO-PRISMLab/AAA-Authenticated-Anonymity-Architecture/actions/workflows/hardhat-test.yaml/badge.svg)](https://github.com/UniBO-PRISMLab/AAA-Authenticated-Anonymity-Architecture/actions/workflows/hardhat-test.yaml)

Smart contracts written in Solidity to implement the AAA protocol.

## Run

Start a development blockchain using

```bash
npx hardat node
```

Deploy the contract using the [deploy script](./scripts/deploy.js)

```bash
npx hardhat compile && npx hardhat run scripts/deploy.js --network localhost
```

## Run with Docker

Build the image

```bash
docker build . -t contracts
```

Run it

```bash
docker run -p 8545:8545 contracts
```

This will start a container that use [anvil](https://getfoundry.sh/anvil/overview/) to provide a local development node (see [start.sh](./scripts/start.sh) for details).

## Test

Run tests with

```bash
npm run test
```

## Documentation

Documentation is automatically generated using [solidity-docgen](https://github.com/OpenZeppelin/solidity-docgen). To generated the documentation use the command

```bash
npm run docs
```

The documentation goes under [./docs](./docs). This can be configured in [hardhat.config.ts](hardhat.config.ts).
