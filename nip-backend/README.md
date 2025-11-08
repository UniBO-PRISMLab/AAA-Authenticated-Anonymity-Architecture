[![go](https://img.shields.io/badge/Go-1.24-00ADD8?logo=Go)](https://go.dev/doc/go1.24)
![swagger](https://img.shields.io/badge/Swagger%20Preview-85EA2D?logo=Swagger&logoColor=black)

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

Currently `docker-compose.yaml` runs 6 instances of `nip-backend` with a shared PostgreSQL databse. To increase the number of instances either switch to [swarm mode](https://docs.docker.com/engine/swarm/) or manually add a `env.instanceN` file and a new `nip-backend` entry in `docker-compose.yaml`.

```yaml
nip-backend-{n}:
  image: nip-backend:latest
  container_name: nip-backend-{n}
  networks:
    - nip-net
  ports:
    - "{port}:8888"
  env_file:
    - .env.instance{n}
```

[Excalidraw Diagram](https://excalidraw.com/#json=J1hfecP489sIF7XvNlq45,GEvU67sRsKGebbV9XD-Cpw)
