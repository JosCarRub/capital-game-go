FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /capital-game ./cmd/capital-game


FROM alpine:latest

COPY --from=builder /capital-game /capital-game

COPY ./data/countries.json /data/countries.json