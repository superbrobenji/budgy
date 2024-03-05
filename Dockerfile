# syntax=docker/dockerfile:1

FROM golang:1.21.6 as dev

RUN apt-get update && apt-get install -y \
    software-properties-common \
    npm
RUN npm install npm@latest -g && \
    npm install n -g && \
    n latest

WORKDIR /app
COPY  . ./
RUN go mod download
RUN make build/docker
EXPOSE 8080
CMD ["/budgy"]
