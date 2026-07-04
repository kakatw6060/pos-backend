# Build stage
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o pos-api ./cmd/main.go

# Final stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/pos-api .
# 建立數據存放目錄
RUN mkdir -p /data
EXPOSE 8080
CMD ["./pos-api"]
