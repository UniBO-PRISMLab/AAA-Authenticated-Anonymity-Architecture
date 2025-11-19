[![go](https://img.shields.io/badge/Go-1.24-00ADD8?logo=Go)](https://go.dev/doc/go1.24)
![swagger](https://img.shields.io/badge/Swagger%20Preview-85EA2D?logo=Swagger&logoColor=black)

NIP (National Identity Provider) server

## Run with Docker

### Setup the database

```bash
echo -n "<name>" > ../db_name.txt
echo -n "<password>" > ../db_password.txt
echo -n "<user>" > ../db_user.txt
```

### Instances Configurations

Create a directory `nip-backend/config` and add as many configuration files as you want (as many instances you need).

The configuration file must be named `.env.instanceN` e.g., `.env.instance1` and should look like this

```env
GIN_MODE="release"
DATABASE_URL="postgres://<user>:<password>@db:5432/<dbname>"
SK="<32 byte base64 encoded secret key>"
ETH_NODE_URL="ws://contracts:8545"
CONTRACT_ADDRESS="0x5FbDB2315678afecb367f032d93F642f64180aa3"
PUBLIC_KEY="<RSA PKCS#8 base64 encoded public key>"
PRIVATE_KEY="<RSA PKCS#8 base64 encoded private key>"
BLOCKCHAIN_PRIVATE_KEY="<blockchain wallet private key from anvil>"
BLOCKCHAIN_ADDRESS="<blockchain wallet address from anvil>"
HTTP_HOST="0.0.0.0"
HTTP_PORT="8888"
```

### Build

Build the `nip-backend` image

```bash
docker build . -t nip-backend
```

### Run

Start a swarm

```bash
docker swarm init
```

And deploy the backend stack

```bash
docker stack deploy -c docker-compose.yaml aaa-backend-stack --detach=false --prune
```

## Sequence Diagrams

[Edit this diagram](https://sequencediagram.org/index.html#initialData=C4S2BsFMAIEF4FAIA4EMBOoDGI0Dtg5wQtIUNtdUDoAiAOVVAHs9VxoBJAE0gLACe0AArpmANxC900ABT1OwgJS1oqAM7QFw8phJUatAMoBbCtADCrYOlRZgmgLSWAFqhB5oAUQAqACWgAGRAAa0hVDWhTcysCW3t1JFhiUkcAPm0ALlRkEAB6cQBGPKk+UGABPNzuBGSSSAAeBsdHLOFOABFalMh06MxLa3iHTPVISG5hF1sxgHE+SFsWPFFmYGYsZnBZJWgEO1BxJhh+wlibA8TwZmZkaAAzZhlPUBNIdQBuaE8AXm+AVxMACNFtBmPdoAB3J7cb7jXg1U6DOKXdKyABMSiykHEZWgAHUYQAlSAAR3+72AE329hARypWkUCG0LTSaQxSiR52G6ky0PQsL4WHQAmQVJqvAOdOOjJ0XKGqPZmLqpEyOLxhIFRn+QJMYHFCD4iLMA25iuxuJoRk6XjwwtFyxJ5Mp1JZ6SVnJNZwVCUy1o60CFIrF1PlKIS6QtePaHVt9rFIFYTop6gNbrZHLDF19McDduDBqzPL6XuR2ZGqaeJ06CEltPpJ1LZoSSR6JZiPpGAHNIMB-TsEEXUe6VZA-Z0gA)

![Sequence Diagram Seed Phrase Flow](docs/sequence-seed.svg)

[Edit this diagram](https://sequencediagram.org/index.html#initialData=C4S2BsFMAIGUEEDC0Bi4D2B3AUNgDgIYBOoAxiIQHbDTzgimT7FkUHXQBEAcgaOpQLhoASQAmkamACe0AApF0ANxASi0ABTcRcgJSdoBAM7Rtc5iQZsOnWAFsW0RAOBECpYCYC0TgBYEQSmgAUQAVAAloABkQAGtIA2M4BxInFzcPIwtWKhpOeEoBaTt0AFcTWEgiFUZEkwKikvK4KpqmbDoGSC8APjMALgI8EAB6JQBGEYJS4F8Ro3dNOQBpXWxC4BhlKtMdfpXoI2B2MRMAM3R1OVKAI3pSaGXIaQ77yAAed68vAYREV66vXsjmc1Aynn6RludjAfwASpBSJcxBo-gAaeSrbDuUBKPgwYGpUGuHFZQk0YngoxAlIU9KkyHAS4EpD7ZbYCQ4kB4zbJEH0zLrdC87bqclpMEM2DLQ7HSinaAXMWIoiQGhPWQaPBEbn46DxWToM6GIzFOxqnUPA2EEBENZcnkwBqUYplCqtLoAxi9Z2u5qVapdfoYADm0ECqKQGJWGNgIgA4txI4hY6s1r6mu7A96euLKQyQ2r4YjkRoVms8wLPN8ej0M26WtnIP0-tBVUiiGIOoUXZnG20fT2-Vm2v0lFUQGdZEYQCHBMBSqru40GwGB70629+vBSIwjCYQ25qJAxNAOZAHXr6-6PYxsEA)

![Sequence Diagram Sac Flow](docs/sequence-sac.svg)

[Edit this diagram](https://sequencediagram.org/index.html#initialData=C4S2BsFMAIAUEEDC0Bi4D2B3AUNgDgIYBOoAxiIQHbDTzgimT7FkUHXQBEAcgaOpQLhoASQAmkamACecIugBuICUWgAKbiNgBKTtAIBnaJtjMSDNh06wArgCN6paAGVIRJYz2G49xy7ceTNh0DJAAtAB8JgBcBHggAPQKAIwJBDbAABYJhKTYJmGRESGM0QiIwY7hEbYODP7uodFwSNgEpKAKfDC1fq6NjNi99f2BkTFxiSlpGdkKbiAAZtJhuflahRE1viMBTfNESyCQYkM7TqOhm8VVzfCkjAZGAOZE7MAn2BLtnd0+dRc9owgA)

![Sequence Diagram Sac Flow](docs/sequence-pac.svg)
