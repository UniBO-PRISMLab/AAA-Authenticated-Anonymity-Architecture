[![go](https://img.shields.io/badge/Go-1.23.1-00ADD8?logo=Go)](https://www.scala-lang.org/download/2.12.18.html)

NIP (National Identity Provider) server

## Run with Docker

Build the `nip-backend` image

```bash
docker build . -t nip-backend
```

Run it with

```bash
docker compose up
```

or

```bash
task up
```

## Scale Instances

Currently `docker-compose.yaml` runs 2 instances of `nip-backend`. To increase the number of instances either switch to [swarm mode](https://docs.docker.com/engine/swarm/) or manually add a `env.instanceN` file and add a new `nip-backend` service in `docker-compose.yaml`.
