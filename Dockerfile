FROM golang:1.22 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o rinha-de-backend-2024-q1 ./cmd/main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/rinha-de-backend-2024-q1 /usr/local/bin/rinha-de-backend-2024-q1

ENV DATABASE_URL=postgres://postgres:postgres@db/postgres

CMD ["rinha-de-backend-2024-q1"]
