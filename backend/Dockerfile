# syntax=docker/dockerfile:1
FROM golang:1.20 AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/server

FROM gcr.io/distroless/static
COPY --from=builder /app/server /app/server
ENTRYPOINT ["/app/server"]
