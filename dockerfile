FROM golang:1.18-alpine

MAINTAINER Nathan Ferreira

ADD . /app

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

ENV HOST=127.0.0.1
ENV PORT=9000
ENV DB_NAME=family-tree-db
ENV DB_HOST=127.0.0.1
ENV DB_PORT=5432
ENV DB_PASSWORD=postgres 
ENV DB_USER=postgres

EXPOSE 9000

CMD ["go", "run", "cmd/api/main.go"]