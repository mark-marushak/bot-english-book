FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /app gofile

CMD ["./gofile"]