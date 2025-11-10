# AAA

Smart contract in solidity for the Authenticated Anonimity Architecture (AAA).

## How to run locally

Start a development blockchain using

```bash
npx hardat node
```

Deploy the contract using the [deploy script](./scripts/deploy.js)

```bash
npx hardhat compile && npx hardhat run scripts/deploy.js --network localhost
```

## Test it

Run tests with

```bash
npm run test
```

This will also produce the gas report under [../docs](../docs/).

The script [seedPhrase.js](./scripts/seedPhrase.js) is a manual test script that request the initialization of the seed generation protocol.

## Documentation

Documentation is automatically generated using [solidity-docgen](https://github.com/OpenZeppelin/solidity-docgen). To generated the documentation use the command.

```bash
npm run docs
```

The documentation goes under [./docs](./docs). This can be configured in [hardhat.config.ts](hardhat.config.ts).
