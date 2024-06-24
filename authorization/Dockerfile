FROM golang:1.21.4 as builder

WORKDIR /auth

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/app/main.go -o ./docs
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/app/main.go

FROM alpine:latest

WORKDIR /smartri_auth

COPY --from=builder /auth/main .
COPY --from=builder /auth/docs ./docs
COPY --from=builder /auth/config/config.yaml ./config/config.yaml

ENTRYPOINT ["./main"]
