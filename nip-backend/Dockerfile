FROM golang:1.24-alpine

WORKDIR /nip-backend

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o nip-backend .

EXPOSE 8080

CMD ["./nip-backend"]