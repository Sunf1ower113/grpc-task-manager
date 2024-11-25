FROM golang:1.22-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o grpc-task-manager ./cmd/server

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/grpc-task-manager /app/grpc-task-manager

RUN mkdir -p /var/log

EXPOSE 50051

CMD ["/app/grpc-task-manager"]
