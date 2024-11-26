FROM golang:1.23

WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o main ./cmd

CMD ["/app/main"]