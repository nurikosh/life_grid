FROM golang:1.26.1-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go test ./...

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

FROM alpine:latest

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /home/appuser
COPY --from=builder /app/main .
COPY --from=builder /app/migrations ./migrations
RUN chown appuser:appgroup ./main

USER appuser

EXPOSE 8080

CMD ["./main"]
