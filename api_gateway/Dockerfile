FROM golang:1.21.4 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/app/main.go

FROM alpine:latest

WORKDIR /smartri_api_gateway

COPY --from=builder /app/main .
COPY --from=builder /app/static ./static
COPY --from=builder /app/config/config.yaml ./config/config.yaml

ENTRYPOINT ["./main"]
