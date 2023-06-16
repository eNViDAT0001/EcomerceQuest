#build stage
FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

#Run stage
FROM alpine:latest
WORKDIR /app/main
COPY --from=builder /app/main .
COPY config /app/main/config

EXPOSE 8082
CMD ["./main"]
