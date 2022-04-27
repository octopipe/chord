FROM golang:1.16-alpine

WORKDIR /app

RUN go mod init sampleserver
RUN go get github.com/octopipe/dht

COPY test/sampleserver/*.go ./

RUN go build -o /server

EXPOSE 8080

CMD [ "/server" ]
