FROM golang:alpine AS builder

WORKDIR /server

COPY . .

RUN go build -o restful_server

FROM alpine:3.17

WORKDIR /server

COPY --from=builder /server .

CMD ./restful_server
