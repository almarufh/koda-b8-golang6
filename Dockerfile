FROM golang:alpine AS builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o gorutine-controll main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/ .

CMD ["/app/gorutine-controll"]