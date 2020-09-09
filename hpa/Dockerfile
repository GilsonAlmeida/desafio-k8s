FROM golang:1.7.0-alpine AS builder

COPY ./go-workspace/src ./


RUN go build main.go



ENTRYPOINT ["./"]