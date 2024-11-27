FROM golang:1.22-alpine as builder

RUN apk --no-cache add git protobuf protobuf-dev curl bash

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest && \
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest && \
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

ENV PATH="/root/go/bin:$PATH"

WORKDIR /app

RUN mkdir -p proto/google/api proto/protoc-gen-openapiv2/options && \
    curl -o proto/google/api/annotations.proto https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto && \
    curl -o proto/google/api/http.proto https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto && \
    curl -o proto/protoc-gen-openapiv2/options/annotations.proto https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/main/protoc-gen-openapiv2/options/annotations.proto && \
    curl -o proto/protoc-gen-openapiv2/options/openapiv2.proto https://raw.githubusercontent.com/grpc-ecosystem/grpc-gateway/main/protoc-gen-openapiv2/options/openapiv2.proto

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN mkdir -p docs
RUN protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --grpc-gateway_out=paths=source_relative:. --openapiv2_out=./docs -I . -I ./proto proto/task.proto

RUN go build -o grpc-task-manager ./cmd/server

FROM alpine:latest

RUN apk --no-cache add ca-certificates && adduser -D -g '' appuser

WORKDIR /app

COPY --from=builder /app/grpc-task-manager /app/grpc-task-manager
COPY --from=builder /app/docs /app/docs
COPY --from=builder /app/proto /app/proto
COPY --from=builder /usr/include/google/protobuf /app/google-protobuf

RUN chown -R appuser /app && chmod +x /app/grpc-task-manager

USER appuser

EXPOSE 50051

CMD ["/app/grpc-task-manager"]
