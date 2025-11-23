[![go](https://img.shields.io/badge/Go-1.24-00ADD8?logo=Go)](https://go.dev/doc/go1.24)
![swagger](https://img.shields.io/badge/Swagger%20Preview-85EA2D?logo=Swagger&logoColor=black)

NIP (National Identity Provider) server

## Run with Docker

### Setup the database

Only need to specifiy the database name, a username and a password. Place them in `db_` file in the root directory.

```bash
echo -n "<name>" > ../db_name.txt
echo -n "<password>" > ../db_password.txt
echo -n "<user>" > ../db_user.txt
```

The docker-compose file will read them and inject them as environment variables in the database container.

### Instances Configurations

Create a directory `/nip-backend/config` and use the script `/gen.sh` (`./gen.sh <n>`) to generate `n` configuration files that will be placed under `/nip-backend/configs`. The script reads the files `/db_*.txt `, generate RSA keypairs using openssl and extract pairs (address, private key) from `/accounts.txt` producing `.env.instanceN`.

### Run

Build the images

```zsh
docker compose build
```

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
