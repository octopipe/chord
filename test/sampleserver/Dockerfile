FROM golang:1.16-alpine

WORKDIR /app

ENV HOST 0.0.0.0
ENV PORT 8080
ENV PARENT 0.0.0.0:8080

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build ./test/sampleserver/*.go

EXPOSE 8080

CMD [ "./main --parent ${PARENT}" ]
