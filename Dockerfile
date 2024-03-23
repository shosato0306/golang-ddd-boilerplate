FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /myapp ./cmd/myapp/main.go

FROM alpine:latest

WORKDIR /

COPY --from=builder /myapp /bin/myapp

ENTRYPOINT ["/bin/myapp"]
