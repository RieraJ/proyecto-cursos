FROM golang:1.23.4-alpine3.21

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]