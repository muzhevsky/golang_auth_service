FROM golang:1.21.4

WORKDIR /app

COPY . .

RUN go build ./cmd/app/main.go

EXPOSE ${AUTH_PORT}

ENTRYPOINT ["./main"]
