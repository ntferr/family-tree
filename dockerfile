FROM golang:latest

ADD . /app

WORKDIR /app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go", "run", "cmd/api/main.go"]

EXPOSE 9090