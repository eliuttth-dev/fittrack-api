FROM golang:1.22.4-alpine3.20

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY . .

RUN go build -o main ./src

EXPOSE 8080

CMD ["./main"]