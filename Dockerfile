FROM golang:latest as builder 
ENV GOPATH /golang
WORKDIR /
COPY ./ /
RUN go build main.go

FROM debian:stretch
ENV GIN_MODE release
COPY --from=0 /main .
COPY --from=0 ./db /db
EXPOSE 8888

ENTRYPOINT ["./main"]