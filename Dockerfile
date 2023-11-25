FROM golang:1.21.4

EXPOSE 8080

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /golang-mongo-crud ./src

CMD ["/golang-mongo-crud"]
