FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod init lab1 && go mod tidy && go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 6080
CMD ["./main"]