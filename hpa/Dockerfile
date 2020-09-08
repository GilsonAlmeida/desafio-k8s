FROM golang:1.7.0-alpine AS builder

#WORKDIR /build



COPY ./go-workspace/src ./


RUN go build main.go

#WORKDIR /dist

#RUN cp  /build  ./

#FROM alpine:latest
#COPY --chown=0:0 --from=builder /dist/ /

ENTRYPOINT ["./"]