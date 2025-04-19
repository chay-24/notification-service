FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server .

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/server .
COPY ./templates /root/templates/
RUN apk add --no-cache sqlite

EXPOSE 8089

CMD ["./server"]
