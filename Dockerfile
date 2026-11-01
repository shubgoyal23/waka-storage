FROM golang:1.27-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o waka-storage .

FROM alpine:3.24

RUN apk --no-cache add ca-certificates tzdata

ENV TZ=UTC

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/waka-storage .

USER appuser

# Background worker with no HTTP port, so check that the process is alive
HEALTHCHECK --interval=30s --timeout=5s --start-period=10s --retries=3 \
  CMD pgrep waka-storage > /dev/null || exit 1

CMD ["./waka-storage"]
