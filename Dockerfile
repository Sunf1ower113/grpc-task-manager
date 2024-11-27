FROM golang:1.22-alpine as builder

RUN apk --no-cache add git bash protobuf protobuf-dev

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN mkdir -p proto/generated && \
    protoc --go_out=./proto --go-grpc_out=./proto \
           --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative \
           proto/task.proto

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
