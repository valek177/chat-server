FROM golang:1.23.2-alpine3.20 AS builder

COPY . /github.com/valek177/chat-server/source/
WORKDIR /github.com/valek177/chat-server/source/

RUN go mod download
RUN go build -o ./bin/chat_server cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/valek177/chat-server/source/bin/chat_server .

CMD ["./chat_server"]
