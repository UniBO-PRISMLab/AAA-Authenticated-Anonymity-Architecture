[![go](https://img.shields.io/badge/Go-1.23.1-00ADD8?logo=Go)](https://www.scala-lang.org/download/2.12.18.html)

NIP (National Identity Provider) server

## Run with Docker

Whole system (server + database)

```bash
docker compose up
```

Server only

```bash
docker compose up nip
```

## Run in plain mode

```bash
go run main.go
```
