FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY .env .env
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o nip-backend .

FROM gcr.io/distroless/static:nonroot
COPY --from=builder /app/nip-backend /nip-backend
COPY --from=builder /app/.env .env
USER nonroot:nonroot
ENTRYPOINT ["/nip-backend"]
