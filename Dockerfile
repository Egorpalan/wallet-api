FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o wallet-api ./cmd/wallet-api

FROM alpine:latest
COPY --from=builder /app/wallet-api /wallet-api
COPY config.env .
CMD ["./wallet-api"]