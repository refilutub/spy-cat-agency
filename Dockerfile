FROM golang:1.24-alpine
LABEL authors="refilutub"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/api

CMD ["./main"]