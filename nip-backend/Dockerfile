FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY words.txt /usr/share/dict/words

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o nip-backend .

FROM gcr.io/distroless/static:nonroot
WORKDIR /app
COPY --from=builder /app/nip-backend .
COPY --from=builder /usr/share/dict/words /usr/share/dict/words
USER nonroot:nonroot
ENTRYPOINT ["/app/nip-backend"]
