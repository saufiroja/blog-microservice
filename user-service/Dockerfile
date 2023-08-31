FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o user-service .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/user-service .
CMD ["./user-service"]