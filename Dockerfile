# syntax=docker/dockerfile:1

FROM golang:1.21.6 
WORKDIR /cmd
COPY go.mod go.sum makefile ./
RUN go mod download
COPY ./cmd/*.go ./
RUN make build/docker
EXPOSE 8080
CMD ["/budgy"]
