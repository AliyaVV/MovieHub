# ---------- BUILD ----------
FROM golang:1.25.1 AS builder

WORKDIR /moviehub


COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o moviehub main.go


# ---------- RUN ----------
FROM alpine:latest

WORKDIR /moviehub

COPY --from=builder /moviehub/moviehub .
COPY .env .
COPY --from=builder /moviehub/internal/storage/postgre/migrations ./internal/storage/postgre/migrations

EXPOSE 8080

CMD ["./moviehub"]