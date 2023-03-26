FROM golang:1.18.2-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

#Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8082
CMD ["/app/main"]