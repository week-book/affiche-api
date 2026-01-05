FROM golang:1.23-alpine AS builder

WORKDIR /build

RUN apk add --no-cache curl

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz \
	| tar xvz

COPY go.mod go.sum ./
RUN go mod download

COPY migrations /migrations
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/api

FROM alpine:3.19-alpine

WORKDIR /app

COPY --from=builder /build/migrate /usr/local/bin/migrate

COPY --from=builder /build/app /app/app

COPY --from=builder /build/migrations /migrations

EXPOSE 8080

CMD ["./app"]

