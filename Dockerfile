FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o main .


FROM alpine:3.20
COPY --from=builder /app/internal/config/config.yaml /config.yaml
COPY --from=builder /app/main /main

EXPOSE 8000
CMD ["/main"]