FROM golang:1.21-alpine AS builder

WORKDIR /go/src/github.com/IlyaZayats/auth
COPY . .

RUN go build -o ./bin/auth ./cmd/auth

FROM alpine:latest AS runner

COPY --from=builder /go/src/github.com/IlyaZayats/auth/bin/auth /app/auth

RUN apk -U --no-cache add bash ca-certificates \
    && chmod +x /app/auth

WORKDIR /app
ENTRYPOINT ["/app/auth"]
