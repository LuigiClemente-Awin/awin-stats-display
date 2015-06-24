FROM golang:1.3

COPY . /go
WORKDIR /go/

RUN go build -v src/pong/pong.go

CMD ["./pong"]