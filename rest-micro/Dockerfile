FROM golang:1.14 as builder
WORKDIR $GOPATH/src/micro
COPY ./ .
RUN GOOS=linux GOARCH=386 go build -ldflags="-w -s" -v
RUN cp microservice /

FROM alpine:latest
COPY --from=builder /microservice /
CMD ["/microservice"]