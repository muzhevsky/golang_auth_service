FROM golang:1.21.4 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/app/main.go

EXPOSE ${AUTH_PORT}

FROM alpine:latest

WORKDIR /root

COPY --from=builder /app/main .
COPY --from=builder /app/config/config.toml ./config/

ENTRYPOINT ["./main"]
