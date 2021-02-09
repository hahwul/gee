# BUILDER
FROM golang:latest AS builder
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go build -o gee

# RUNNING
FROM debian:buster
RUN mkdir /app
WORKDIR /app/
COPY --from=builder /go/src/app/gee /app/gee
CMD ["/app/gee"]
