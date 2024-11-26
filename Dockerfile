FROM golang:1.22-alpine as builder

RUN apk --no-cache add git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o grpc-task-manager ./cmd/server

FROM alpine:latest

RUN apk --no-cache add ca-certificates && \
    adduser -D -g '' appuser

WORKDIR /app

COPY --from=builder /app/grpc-task-manager /app/grpc-task-manager

RUN mkdir -p /app/logs && chown -R appuser /app/logs

USER appuser

EXPOSE 50051

CMD ["/app/grpc-task-manager"]
