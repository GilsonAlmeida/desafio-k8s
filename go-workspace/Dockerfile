FROM golang:1.7.0-alpine AS builder

WORKDIR /build



COPY ./src/hello ./


RUN go build hello.go

WORKDIR /dist

RUN cp  /build/hello  ./hello

FROM alpine:latest
COPY --chown=0:0 --from=builder /dist/hello /hello

ENTRYPOINT ["./hello"]