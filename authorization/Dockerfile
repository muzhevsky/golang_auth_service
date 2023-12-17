FROM golang:1.21.4

WORKDIR /app

COPY . .

RUN go build ./cmd/app/main.go

EXPOSE 8080

CMD ["./main"]
